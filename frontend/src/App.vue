<template>
  <div class="app-container">
    <hsidebar
      class="sidebar"
      :visible="sidebarOpen"
      :active-view="currentView"
      @open-compose="showComposeView"
      @open-inbox="showSectionView('inbox')"
      @open-sent="showSectionView('sent')"
      @open-starred="showSectionView('starred')"
      @open-drafts="showSectionView('drafts')"
      @open-spam="showSectionView('spam')"
      @open-trash="showSectionView('trash')"
      @open-all="showSectionView('all')"
      @open-important="showSectionView('important')"
    />

    <div v-if="isLoading" class="loader-overlay">
      <div class="loader">
        <i class="fa-solid fa-spinner fa-spin"></i>
        <span>Loading...</span>
      </div>
    </div>

    <div :class="['main-content', { 'with-sidebar': sidebarOpen }]">
      <TopBar class="topbar" @toggle-sidebar="toggleSidebar" />

      <div id="app">
        <ToolBar />
        <Compose
          v-if="currentView === 'compose'"
          @close="showInboxView"
        />
        <EmailView
          v-else-if="currentView === 'emailview'"
          :email="emailToView"
          @close="closeEmailView"
        />
        <EmailList
          v-else
          ref="emailListComponent"
          :section="sectionToView"
          @viewEmail="openEmailView"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import { ReadEmail } from '../wailsjs/go/main/App';
import { EventsOn } from '../wailsjs/runtime/runtime'; // Import EventsOn
import Compose from './components/Compose.vue';
import EmailList from './components/EmailList.vue';
import TopBar from './components/TopBar.vue';
import hsidebar from './components/hsidebar.vue';
import ToolBar from './components/ToolBar.vue';
import EmailView from './components/EmailView.vue';
import Home from './components/Home.vue';

const sidebarOpen = ref(true);
const isLoading = ref(false); // Parent now controls global loading

const currentView = ref('inbox');
const sectionToView = ref('INBOX'); // Holds the IMAP folder name
const previousView = ref('inbox'); // To remember where to go back to
const emailToView = ref(null);
const emailListComponent = ref(null); // A ref to the EmailList component instance

const toggleSidebar = () => {
  sidebarOpen.value = !sidebarOpen.value;
};

// Close sidebar on ESC key
const onKey = (e) => {
  if (e.key === 'Escape') {
    sidebarOpen.value = false;
  }
};


// A map to translate simple names to Gmail's IMAP folder names
const sectionMap = new Map([
  ['inbox', 'INBOX'],
  ['sent', '[Gmail]/Sent Mail'],
  ['starred', '[Gmail]/Starred'],
  ['drafts', '[Gmail]/Drafts'],
  ['spam', '[Gmail]/Spam'],
  ['trash', '[Gmail]/Trash'],
  ['allmail', '[Gmail]/All Mail'],
  ['important', '[Gmail]/Important'],
]);

const showComposeView = () => {
  emailToView.value = null;
  currentView.value = 'compose';
};

const showInboxView = () => {
  showSectionView('inbox');
};

const showSectionView = (sectionKey) => {
  emailToView.value = null;
  sectionToView.value = sectionMap.get(sectionKey) || 'INBOX';
  currentView.value = sectionKey; // All list views are technically the 'inbox' type view
};

const openEmailView = async (email) => {
  isLoading.value = true;
  try {
    const fullEmailData = await ReadEmail(email.seqNum, sectionToView.value); // Pass section for context
    emailToView.value = fullEmailData;
    previousView.value = currentView.value; // Remember the view we came from
    currentView.value = 'emailview';
  } catch (err) {
    console.error(err);
  } finally {
    isLoading.value = false;
  }
};

const closeEmailView = () => {
  emailToView.value = null;
  currentView.value = previousView.value; // Go back to the previous view
  // No need to reload here, the watcher in EmailList handles it if the section changed
};

onMounted(() => {
  // Solve the race condition: Wait for the Go backend to be ready
  EventsOn("backend:ready", () => {
    console.log("Backend is ready, loading initial emails.");
    if (emailListComponent.value) {
      emailListComponent.value.loadEmails();
    }
  });
});
</script>

<style>
/* Reset */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  margin: 0;
  font-family: 'Segoe UI', sans-serif;
  background-color: #fff;
}

/* Style the new loader overlay */
.loader-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(255, 255, 255, 0.8); /* Semi-transparent white */
  z-index: 10;
  display: flex;
  align-items: center;
  justify-content: center;
}

.loader {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 30px;
  color: #666;
  font-weight: 500;
  font-size: 1.2em;
}

.loader i {
  margin-right: 15px;
  font-size: 1.5em;
}

/* App layout */
.app-container {
  height: 100vh;
  width: 100%;
  position: relative;
  overflow-x: hidden;
}

.topbar {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 56px;
  z-index: 100;
}

.sidebar {
  position: fixed;
  top: 0;
  left: 0;
  height: 96vh;
  width: 300px;
  z-index: 1000;
  margin-top: 0;
  
  transition: margin-left 0.28ms ease-in, width 0.28ms ease-in-out;
}

.main-content {
  width: 100%;
  height: calc(100vh - 56px);
  padding: 5px 15px;
  transition: margin-left 0.28s ease, width 0.28s ease;
  margin-top: 56px;
  display: flex;
  flex-direction: column;
}

.main-content.with-sidebar {
  margin-left: 300px;
  width: calc(100% - 300px);
}

#app {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.compose-container {
  flex: 1;
  width: 100%;
  padding: 16px;
  box-sizing: border-box;
}

.welcome-screen {
  flex: 1;
  text-align: left;
  color: #555;
  z-index: 100;
  padding: 12px 0;
  display: flex;
  align-items: center;
  justify-content: flex-start;
}

/* Sidebar helper layout (used by components) */
.menu-area {
  position: relative;
  z-index: 10;
}

.menu-list-wrapper {
  position: absolute;
  top: 0;
  left: 0;
  width: 300px;
  background: #f5f5f5;
  border-right: 1px solid #e0e0e0;
  height: 100%;
  padding-top: 56px;
}

/* Small helpers */
.content-area {
  width: 100%;
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #fff;
}

#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
