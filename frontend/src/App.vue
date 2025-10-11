<template>
  <div>
    <TopBar />
    <div id="app">
      <Compose v-if="isComposeOpen" @close="closeCompose" />
      
      <EmailView v-else-if="selectedEmail" :email="selectedEmail" @close="closeEmailView" />
      
      
      <Inbox v-else @view-email="viewEmail" />
      
  
      <button v-if="!isComposeOpen && !selectedEmail" @click="openCompose" class="floating-compose-btn">
        <i class="fa-solid fa-pen"></i>
        <span>Compose</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import Compose from './components/Compose.vue';
import Inbox from './components/EmailList.vue';
import TopBar from './components/TopBar.vue';

import EmailView from './components/EmailView.vue'; 

const isComposeOpen = ref(false);
const selectedEmail = ref(null);

const openCompose = () => {
  selectedEmail.value = null; // Clear emailview for compose
  isComposeOpen.value = true;
};
const closeCompose = () => (isComposeOpen.value = false);

const viewEmail = (email) => {
  isComposeOpen.value = false; // Close compose window if it's open
  selectedEmail.value = email;
};

const closeEmailView = () => {
  selectedEmail.value = null;
};
</script>

<style>

body {
  margin: 0;
  font-family: 'Segoe UI', sans-serif;
  background-color: #f6f8fa;
  overflow: hidden; /*NO scrolling */
}

#app {
  height: 100vh;
  width: 100vw;
  position: relative; /* floating button */
}

.floating-compose-btn {
  position: fixed;
  bottom: 24px;
  right: 24px;
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

.floating-compose-btn:hover {
  box-shadow: 0 6px 20px rgba(0,0,0,.2);
  background-color: #f8f9fa;
}

.floating-compose-btn i {
  font-size: 18px;
}
</style>