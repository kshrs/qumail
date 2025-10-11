<template>
  <div class="app-container">
    <ToolBar />
    <Inbox />
    <!-- Sidebar -->
    <hsidebar class="sidebar" :visible="sidebarOpen" @open-compose="openCompose" />


    <!-- Main content -->
    <div :class="['main-content', { 'with-sidebar': sidebarOpen }]">
      <TopBar class="topbar" @toggle-sidebar="toggleSidebar" />

      <div id="app">
        <div v-if="isComposeOpen" class="compose-container">
          <Compose @close="closeCompose" />
        </div>

        <div v-if="!isComposeOpen" class="welcome-screen">
        </div>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import Compose from './components/Compose.vue';
import Inbox from './components/EmailList.vue';
import TopBar from './components/TopBar.vue';
import hsidebar from './components/hsidebar.vue';
import ToolBar from './components/ToolBar.vue';
import EmailView from './components/EmailView.vue'; 


const sidebarOpen = ref(false);
const isComposeOpen = ref(false);

function toggleSidebar() {
  sidebarOpen.value = !sidebarOpen.value;
}


const selectedEmail = ref(null);

const openCompose = () => {
  selectedEmail.value = null; // Clear emailview for compose
  isComposeOpen.value = true;
  sidebarOpen.value = true;
};
const closeCompose = () => {
  isComposeOpen.value = false;
};

const viewEmail = (email) => {
  isComposeOpen.value = false; // Close compose window if it's open
  selectedEmail.value = email;
};

const closeEmailView = () => {
  selectedEmail.value = null;
};


// Close sidebar on ESC key
const onKey = (e) => {
  if (e.key === 'Escape') {
    sidebarOpen.value = false;
    if (isComposeOpen.value) {
      closeCompose();
    }
  }
};

onMounted(() => {
  window.addEventListener('keydown', onKey);
});

onUnmounted(() => {
  window.removeEventListener('keydown', onKey);
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
  z-index: 1001;
}

.sidebar {
  position: fixed;
  top: 0;
  left: 0;
  height: 100vh;
  width: 300px;
  z-index: 1000;
  
  transition: margin-left 0.28ms ease-in, width 0.28ms ease-in-out;
}

.main-content {
  width: 100%;
  height: calc(100vh - 56px);
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
  box-shadow: 2px 0 8px rgba(0, 0, 0, 0.08);
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
  gap: 15px;
  background-color: #ffffff;
  color: #3c4043;
  border: 1px solid #dadce0;
  border-radius: 28px;
  cursor: pointer;
  font-size: 15px;
  font-weight: 600;
  box-shadow: 0 4px 12px rgba(0,0,0,.15);
  transition: all 0.2s ease-in-out;
  padding: 16px 24px;
  z-index: 999;
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
