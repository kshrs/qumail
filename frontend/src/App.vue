<template>
  <div class="app-container">
    <TopBar v-if="loggedIn" class="topbar" @toggle-sidebar="toggleSidebar" />
    <ToolBar v-if="loggedIn" />
    <!-- Sidebar -->
    <hsidebar v-if="loggedIn" class="sidebar" :visible="sidebarOpen" @open-compose="openCompose" />
    <Login_Page v-if="!loggedIn" @login="handleLogin" />
    <!-- Main content -->
    <div v-if="loggedIn" :class="['main-content', { 'with-sidebar': sidebarOpen }]">

      <div id="app">
        <div v-if="isComposeOpen" class="compose-container">
          <Compose @close="closeCompose" />
        </div>

        <div v-if="!isComposeOpen" class="welcome-screen">
          <!-- show inbox at top -->
          <Inbox />
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
import Login_Page from './components/Login_Page.vue';

const sidebarOpen = ref(false);
const loggedIn = ref(false); // show login page first

function toggleSidebar() {
  sidebarOpen.value = !sidebarOpen.value;
}


const isComposeOpen = ref(false);

const openCompose = () => {
  isComposeOpen.value = true;
  sidebarOpen.value = true;
};

// handle login event from Login_Page
function handleLogin(payload) {
  const { email, password } = payload || {};
  // dummy check; replace with server-side DB check in production
  if (email === 'kabb@projectbase.com' && String(password) === '1234') {
    loggedIn.value = true;
  } else {
    alert('Invalid credentials. Use kabb@projectbase.com / 1234');
  }
}

const closeCompose = () => {
  isComposeOpen.value = false;
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
  user-select: none;
}

body {
  margin: 0;
  font-family: 'Segoe UI', sans-serif;
  background-color: #fff;
}

/* App layout */
.app-container {
  margin-top:50px;
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
