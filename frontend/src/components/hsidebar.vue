<template>
<transition name="slide">
    <div v-if="visible" class="sidebar">
      <div class="sidebar-top">
        <button class="sidebar-compose" @click="$emit('open-compose')">
          <i class="fa-solid fa-pen"></i>
          <span>Compose</span>
        </button>
      </div>
      <HMenuList 
      :active-view="activeView"
      @open-inbox="$emit('open-inbox')"
      @open-sent="$emit('open-sent')"
      @open-spam="$emit('open-spam')"
      @open-drafts="$emit('open-drafts')"
      @open-trash="$emit('open-trash')"
      @open-allmail="$emit('open-allmail')"
      @open-starred="$emit('open-starred')"
      @open-important="$emit('open-important')"
      />
    </div>
  </transition>

</template>

<script setup>
import HMenuList from './HMenuList.vue';
defineProps({ visible: Boolean, activeView: String });
defineEmits(['open-compose', 'open-inbox', 'open-sent', 'open-spam', 'open-drafts', 'open-trash', 'open-allmail', 'open-starred', 'open-important']);

</script>

<style scoped>
.sidebar {
  width: 300px;
  background: #1E1E1E; /* Zen: UI/Surface Background */
  color: #E1E1E1; /* Zen: Primary Text */
  padding: 32px 24px 24px 24px;
  height: calc(100vh - 56px);
  margin-top: 56px;
  box-shadow: none; /* Removed for a flatter, modern look */
  border-right: 1px solid #2a2a2a; /* Subtle border for separation */
  border-radius: 0 12px 12px 0;
  display: flex;
  flex-direction: column;
  gap: 16px;
  z-index: 100;
}
.sidebar-top {
  padding-top: 8px;
  display: flex;
  align-items: center;
  justify-content: flex-start;
}
.sidebar-compose {
  display: inline-flex;
  gap: 10px;
  align-items: center;
  background: transparent; /* Maintains the outline style */
  color: #d1cfc0; /* Melon accent color */
  background-color: #363636;
  padding: 8px 14px;
  cursor: pointer;
  font-size: 16px;
  font-weight: 600;
  border-radius: 25px;
  border: 2px solid #363636; /* Melon accent border */
  transition: background-color 0.2s ease, color 0.2s ease;
  outline: none;
}

.sidebar-compose:hover {
  background: #1f1f1f; /* Fill with melon accent color on hover */
  color: #fff0df; /* Use a dark color for the text for contrast */
}

.sidebar-compose i { font-size: 14px; }
/* When sidebar is present it covers left area; slide transition in App.vue will animate transform */
</style>
