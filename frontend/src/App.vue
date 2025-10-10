<template>
  <TopBar />
  <div id="app">
    <div v-if="isComposeOpen">
      <Compose @close="closeCompose" />
    </div>
    <div v-else class="welcome-screen">
      <button class="compose-btn" @click="openCompose">
        <i class="fa-solid fa-pen"></i>
        <span>Compose</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import Compose from './components/Compose.vue';
import TopBar from './components/TopBar.vue';

const isComposeOpen = ref(false);
const openCompose = () => (isComposeOpen.value = true);
const closeCompose = () => (isComposeOpen.value = false);
</script>

<style>
body {
  margin: 0;
  font-family: 'Segoe UI', sans-serif;
  background-color: #f6f8fa;
}

#app {
  height: 100vh;
  width: 100vw;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.welcome-screen {
  text-align: center;
  color: #555;
}

.compose-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 15px;
  background-color: #fffdfd;
  color: #16498c;
  border: none;
  border-radius: 28px;
  cursor: pointer;
  font-size: 15px;
  font-weight: 600;
  box-shadow: 0 1px 3px rgba(0,0,0,.1), 0 1px 2px rgba(0,0,0,.06);
  transition: all 0.3s ease-in-out;
  padding: 15px 30px;
  margin-bottom: 20px;
}

.compose-btn:hover {
  background-color: #c7cace;
}

.compose-btn i {
  font-size: 18px;
}
</style>
