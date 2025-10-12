<template>
  <div class="top-bar">
    <button class="hamburger" @click="$emit('toggle-sidebar')">☰</button>
    <div class="top-btns">
      <button class="top-btn">Home</button>
      <button class="top-btn">View</button>
      <button class="top-btn">Help</button>
    </div>
  </div>
</template>

<script setup>
// Emits 'toggle-sidebar' when hamburger is clicked
</script>

<style scoped>
.top-bar {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 500;
  height: 5vh;
  display: flex;
  align-items: center;
  background: #007bff;
  padding: 12px 24px;
  color: white;
  box-shadow: 0 2px 6px rgba(0,0,0,0.12);
}
.hamburger {
  background: transparent;
  color: white;
  border: none;
  font-size: 24px;
  cursor: pointer;
  margin-right: 24px;
}
.top-btns {
  display: flex;
  align-items: center;
}
.top-btn {
  background: transparent;
  color: white;
  border: none;
  margin-right: 20px;
  font-size: 16px;
  cursor: pointer;
  font-weight: bold;
}
.top-btn:last-child {
  margin-right: 0;
}
.top-btn:hover {
  text-decoration: underline;
}
</style>
