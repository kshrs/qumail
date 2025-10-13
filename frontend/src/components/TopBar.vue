<template>
  <div class="top-bar">
    <button class="hamburger" @click="$emit('toggle-sidebar')">☰</button>
    <div class="top-btns">
      <button class="top-btn">Home</button>
      <button class="top-btn">View</button>
      <button class="top-btn">Help</button>
      <div class="plugins-wrapper">
        <button class="top-btn" @click="showPlugins = !showPlugins" ref="pluginsBtn">Plugins</button>
        <div v-if="showPlugins" class="plugins-popover">
          <div class="plugins-actions icons">
            
            <button class="icon-action" title="Call" @click="onCall">
                <div class="material-symbols-outlined">
                    call
                </div>
            </button>
            <button class="icon-action" title="Chat" @click="onChat">
                <div class="material-symbols-outlined">
                    chat
                </div>
            </button>
            <button class="icon-action" title="Add" @click="openAddDialog">
                <div class="material-symbols-outlined">
                    add
                </div>
            </button>
          </div>

          <!-- only action icons shown here (Add / Call / Chat) -->
        </div>
      </div>
    </div>

    <!-- inline add handled inside popover -->
    <!-- Add-plugin modal (shows Plugin 1 and Plugin 2) -->
    <Transition name="fade">
    <div v-if="showAddDialog" class="add-dialog-backdrop" @click="closeAddDialog">
      <div class="add-dialog" @click.stop>
        <h3>Available plugins</h3>
        <div class="available-list">
          <div class="available-item">Ai Draft <button class="small" @click="addPluginFromDialog('Plugin 1')">Add</button></div>
          <div class="available-item">Schedular <button class="small" @click="addPluginFromDialog('Plugin 2')">Add</button></div>
        </div>
        <div style="text-align:right;margin-top:12px;padding-right: 120px;"><button @click="closeAddDialog" class="small">Close</button></div>
      </div>
    </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue';

// Emits 'toggle-sidebar' when hamburger is clicked
const showPlugins = ref(false);
const showAddDialog = ref(false);
const plugins = ref(['Plugin A','Plugin B']);

function openAddDialog() {
  // close the popover and open the modal dialog
  showPlugins.value = false;
  showAddDialog.value = true;
}

function closeAddDialog() {
  showAddDialog.value = false;
}

function addPluginFromDialog(name) {
  if (name && name.trim()) plugins.value.push(name.trim());
  closeAddDialog();
}

function onCall() {
  // placeholder: actual call integration goes here
  alert('Call action clicked (plugin)');
  showPlugins.value = false;
}

function onChat() {
  // placeholder: open chat for plugin
  alert('Chat action clicked (plugin)');
  showPlugins.value = false;
}

// basic outside click handler (component-local)
let onDocClick = (e) => {
  const el = e.target;
  // if click happens outside any open plugin popover, close it
  if (!el.closest || !document) return;
  const pop = document.querySelector('.plugins-popover');
  const btn = document.querySelector('.plugins-wrapper .top-btn');
  if (pop && btn && !pop.contains(el) && !btn.contains(el)) {
    showPlugins.value = false;
  }
};

onMounted(() => document.addEventListener('click', onDocClick));
onBeforeUnmount(() => document.removeEventListener('click', onDocClick));
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-active .confirm-dialog,
.fade-leave-active .confirm-dialog {
  transition: transform 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0.0;
}
.fade-enter-from .confirm-dialog,
.fade-leave-to .confirm-dialog {
  transform: scale(0.95);
}
.top-bar {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 7vh;
  z-index: 500;
  display: flex;
  align-items: center;
  background: #007bff;
  padding: 12px 24px;
  padding-top:10px;
  color: white;
  box-shadow: 0 2px 6px rgba(0,0,0,0.12);
  border-radius: 0;
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

.plugins-wrapper {
  position: relative;
}
.plugins-popover {
  position: absolute;
  left: 0;
  top: 40px;
  width: 220px;
  background: #fff;
  color: #111;
  border-radius: 8px;
  box-shadow: 0 8px 30px rgba(0,0,0,0.12);
  padding: 8px;
  z-index: 1200;
}
.plugins-list { display:flex; flex-direction:column; gap:8px; }
.plugin-item { display:flex; justify-content:space-between; align-items:center; padding:6px 8px; border-radius:6px; }
.plugin-item:hover { background: #f3f4f6 }
.small { background: #2563eb; color:#fff; border:none; padding:6px 8px; border-radius:6px; cursor:pointer }

/* dialog removed: using inline add inside popover */
.add-dialog-backdrop { position: fixed; inset: 0; display:flex; align-items:center; justify-content:center; background: rgba(0,0,0,0.35); z-index: 2000 }
.add-dialog { background: black; border-radius: 10px; padding: 16px; width: 320px; box-shadow: 0 10px 40px rgba(0,0,0,0.2); }
.available-list { display:flex; flex-direction:column; gap:8px; margin-top:8px }
.available-item { display:flex; justify-content:space-between; align-items:center; padding:8px; border-radius:8px }
.available-item:hover { background:#f8fafc ;color:black}
.icon-action{
  width:50px;
  height: 50px;
  background: none;       /* Remove background color */
  border: none;           /* Remove border */
  padding: 0;             /* Remove padding */
  margin: 0;              /* Remove margin (optional) */
  color: inherit;         /* Inherit text color from parent */
  font: inherit;          /* Inherit font style */
  cursor: pointer;        /* Keep pointer cursor */
}
</style>
