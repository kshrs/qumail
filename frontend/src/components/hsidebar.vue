<template>
<transition name="slide">
    <div v-if="visible" class="sidebar">
      <div class="sidebar-top">
        <button class="sidebar-compose" @click="$emit('open-compose')">
          <i class="fa-solid fa-pen"></i>
          <span>Compose</span>
        </button>
      </div>
      <HMenuList />
    </div>
  </transition>

</template>

<script setup>
import { defineProps } from 'vue';
import HMenuList from './HMenuList.vue';
const props = defineProps({ visible: Boolean });
</script>

<style scoped>
.sidebar {
  width: 300px;
  background: #fafafa;
  color: black;
  padding: 32px 24px 24px 24px;
  height: calc(100vh - 56px);
  margin-top: 56px;
  box-shadow: 2px 0 8px rgba(0,0,0,0.08);
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
  background: #fffdfd;
  color: #16498c;
  border: none;
  border-radius: 20px;
  padding: 8px 14px;
  cursor: pointer;
  font-weight: 600;
}
.sidebar-compose i { font-size: 14px; }
/* When sidebar is present it covers left area; slide transition in App.vue will animate transform */
</style>