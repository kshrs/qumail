<template>
  <div id="app">
    <Compose v-if="isComposeOpen" @close="closeCompose" />
    <Inbox v-else />

    <button v-if="!isComposeOpen" @click="openCompose" class="floating-compose-btn">
      <i class="fa-solid fa-pen"></i>
      <span>Compose</span>
    </button>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import Compose from './components/Compose.vue';
import Inbox from './components/EmailList.vue';

const isComposeOpen = ref(false);
const openCompose = () => (isComposeOpen.value = true);
const closeCompose = () => (isComposeOpen.value = false);
</script>

<style>
body {
  margin: 0;
  font-family: 'Segoe UI', sans-serif;
  background-color: #f6f8fa;
  overflow: hidden; /* Prevent the page from scrolling */
}

/* These styles allow your components to fill the screen properly */
#app {
  height: 100vh;
  width: 100vw;
  position: relative; /* Needed for the floating button */
}

/* Styles for the floating compose button */
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
  z-index: 999; /* This ensures the button is on top of the inbox list */
}

.floating-compose-btn:hover {
  box-shadow: 0 6px 20px rgba(0,0,0,.2);
  background-color: #f8f9fa;
}

.floating-compose-btn i {
  font-size: 18px;
}
</style>