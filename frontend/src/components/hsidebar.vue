<template>
<transition name="slide">
    <div v-if="visible" class="sidebar">
      <div class="sidebar-top">
        <button class="sidebar-compose" @click="$emit('open-compose')">
          <i class="fa-solid fa-pen"></i>
          <span>Compose</span>
        </button>
      </div>
      <HMenuList @open-inbox="$emit('open-inbox')" />
    </div>
  </transition>
</template>

<script setup>
import HMenuList from './HMenuList.vue';
defineProps({ visible: Boolean });
defineEmits(['open-compose', 'open-inbox']);

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
  background: #3174f1;
  color: #ffffff;
  padding: 8px 14px;
  cursor: pointer;
  font-size: 16px;
  font-weight: 600;

  border-radius: 25px;
  border: 3px solid #3174f1;
}
.sidebar-compose i { font-size: 14px; }
/* When sidebar is present it covers left area; slide transition in App.vue will animate transform */
</style>
