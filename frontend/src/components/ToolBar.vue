<script lang="ts" setup>
import { ref, reactive, onMounted, onBeforeUnmount } from 'vue'

// --- New Mail Button
function newMail() {
  console.log("New Mail Clicked")
}

// --- Navigation helper
function goTo(href: string) {
  window.location.href = href
}

// --- Delete button confirmation dialog
const showDeleteDialog = ref(false)
function confirmDelete() {
  showDeleteDialog.value = true
}

function cancelDelete() {
  showDeleteDialog.value = false
}

function deleteAction() {
  console.log("Delete action confirmed")
  showDeleteDialog.value = false
}
</script>

<template>
  <main class="top-bar-container">
    <div class="top-bar">
      <div class="feature-buttons">
        <!-- Feature Buttons -->
        <button class="feature-btn report-btn" @click="goTo('#')" title="Report">Report</button>
        <button class="feature-btn reply-btn" @click="goTo('#')" title="Reply">Reply</button>
        <button class="feature-btn reply-all-btn" @click="goTo('#')" title="Reply All">Reply All</button>
        <button class="feature-btn forward-btn" @click="goTo('#')" title="Forward">Forward</button>
        <button class="feature-btn mark-read-btn" @click="goTo('#')" title="Mark as Read">Mark as Read</button>
        <button class="feature-btn starred-btn" @click="goTo('#')" title="Starred">Starred</button>
        <button class="feature-btn important-btn" @click="goTo('#')" title="Important">Important</button>
        <button class="feature-btn spam-btn" @click="goTo('#')" title="Spam">Spam</button>
        <button class="feature-btn delete-btn" @click="confirmDelete" title="Delete">Delete</button>
        <button class="feature-btn settings-btn" @click="goTo('#')" title="Settings">Settings</button>
      </div>

    </div>

    <!-- Delete Confirmation Dialog -->
    <div v-if="showDeleteDialog" class="delete-dialog">
      <p class="delete-dialog-text">Are you sure you want to delete?</p>
      <div class="delete-dialog-buttons">
        <button class="dialog-btn confirm-btn" @click="deleteAction">Yes</button>
        <button class="dialog-btn cancel-btn" @click="cancelDelete">No</button>
      </div>
    </div>
  </main>
</template>

<style scoped>
/* Top Bar Container */
.top-bar-container {
  padding: 100px;
}

/* Top Bar Styles */
.top-bar {
  display: flex;
  align-items: center;
  gap: 20px; /* Ensure even spacing between buttons */
  padding: 10px 12px;
  background: linear-gradient(180deg, #fbfcfd 0%, #f2f6f9 100%);
  border: 1px solid rgba(200, 200, 200, 0.6);
  border-radius: 8px;
}

.feature-buttons {
  display: flex;
  gap: 20px; /* Ensure even spacing between buttons */
}

.feature-btn {
  background: #e2e8f0;
  color: #1a202c;
  border: 1px solid #cbd5e0;
  padding: 6px 12px;
  border-radius: 4px;
  cursor: pointer;
  position: relative;
}

.feature-btn:hover {
  background: #ffffff;
}

.feature-btn::after {
  content: attr(title);
  position: absolute;
  bottom: -30px;
  left: 50%;
  transform: translateX(-50%);
  background: #333;
  color: #fff;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
  white-space: nowrap;
  opacity: 0;
  pointer-events: none;
  transition: opacity 0.2s ease-in-out;
}

.feature-btn:hover::after {
  opacity: 1;
}

/* Ensure only one hover box is displayed */
.feature-btn:hover::before {
  display: none;
}

/* Dropdown Container */
.dropdown-container {
  position: relative;
}

.readable-select {
  min-width: 260px;
  max-width: 420px;
  text-align: left;
  padding: 8px 12px;
  border-radius: 4px;
  border: 1px solid rgba(150, 150, 150, 0.3);
  background: white;
  cursor: pointer;
  white-space: normal; /* allow wrapping */
}

.select-dropdown {
  position: absolute;
  top: 100%;
  left: 0;
  background: white;
  border: 1px solid #d0d7de;
  box-shadow: 0 6px 18px rgba(0, 0, 0, 0.08);
  border-radius: 6px;
  z-index: 1000;
  min-width: 260px;
  max-width: 420px;
  margin-top: 6px;
}

.select-item {
  padding: 8px 12px;
  cursor: pointer;
  color: #111827;
}

.select-item:hover {
  background: #f3f4f6;
}

/* Delete Confirmation Dialog */
.delete-dialog {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: #fff;
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  z-index: 1000;
}

.delete-dialog-text {
  margin-bottom: 20px;
  font-size: 16px;
  color: #333;
}

.delete-dialog-buttons {
  display: flex;
  gap: 10px;
}

.dialog-btn {
  padding: 6px 12px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.confirm-btn {
  background: #f87171;
  color: #fff;
}

.confirm-btn:hover {
  background: #e01b1b;
  color: #fff;
}

.cancel-btn {
  background: #e2e8f0;
  color: #1a202c;
}

.cancel-btn:hover {
  background: #353535;
  color: #c4c7ce;
}
</style>
