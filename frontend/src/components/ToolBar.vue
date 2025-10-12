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

        <button class="feature-btn report-btn" title="Report">
          <span class="material-symbols-outlined">report</span>
          <span class="label">Report</span>
        </button>

        <button class="feature-btn reply-btn" title="Reply">
          <span class="material-symbols-outlined">reply</span>
          <span class="label">Reply</span>
        </button>

        <button class="feature-btn reply-all-btn" title="Reply All">
          <span class="material-symbols-outlined">reply_all</span>
          <span class="label">Reply All</span>
        </button>

        <button class="feature-btn forward-btn" title="Forward">
          <span class="material-symbols-outlined">forward</span>
          <span class="label">Forward</span>
        </button>

        <button class="feature-btn mark-read-btn" title="Mark as Read">
          <span class="material-symbols-outlined">mark_as_unread</span>
          <span class="label">Mark Read</span>
        </button>
        
        <button class="feature-btn starred-btn" title="Starred">
          <i class="fa-regular fa-star"></i>
          <span class="label">Star</span>
        </button>

        <button class="feature-btn spam-btn" title="Spam">
          <span class="material-symbols-outlined">dangerous</span>
          <span class="label">Junk</span>
        </button>

        <button class="feature-btn delete-btn" title="Delete">
          <span class="material-symbols-outlined">delete</span>
          <span class="label">Delete</span>
        </button>

        <div class="separator"></div>

        <button class="feature-btn settings-btn" title="Settings">
          <span class="material-symbols-outlined">settings</span>
          <span class="label">Settings</span>
        </button>

      </div>
    </div>

    <div v-if="showDeleteDialog" class="delete-dialog">
        </div>
  </main>
</template>

<style scoped>
/* Main container for the top bar */
.top-bar-container {
  -webkit-user-select: none; /* Prevents text selection */
  user-select: none;
}

/* The bar itself */
.top-bar {
  display: flex;
  align-items: center;
  padding: 4px;
  background-color: #f5f5f5; /* Outlook's light grey background */
  border-bottom: 1px solid #e1e1e1;
  width: 100%;
}

.feature-buttons {
  display: flex;
  align-items: center;
}

/* Base style for all feature buttons */
.feature-btn {
  display: flex;
  flex-direction: column; /* Stack icon on top of label */
  align-items: center;
  justify-content: center;
  background: transparent;
  border: none;
  border-radius: 4px; /* Subtle rounded corners on hover */
  cursor: pointer;
  padding: 4px 8px;
  min-width: 60px; /* Ensure buttons have a uniform width */
  transition: background-color 0.2s ease-in-out;
}

/* Hover state for all buttons */
.feature-btn:hover {
  background-color: #e1e1e1;
}

/* --- Icon and Label Base Styles --- */
.feature-btn .material-symbols-outlined,
.feature-btn .fa-regular {
  font-size: 22px;
  color: #5f6368; /* Default neutral icon color */
  transition: color 0.2s ease-in-out;
}

.feature-btn .label {
  font-size: 11px;
  margin-top: 2px;
  color: #5f6368; /* Default neutral text color */
  transition: color 0.2s ease-in-out;
}

/* --- Specific Hover Colors --- */

/* Blue for primary actions: Reply, Reply All, Forward */
.reply-btn:hover .material-symbols-outlined,
.reply-btn:hover .label,
.reply-all-btn:hover .material-symbols-outlined,
.reply-all-btn:hover .label,
.forward-btn:hover .material-symbols-outlined,
.forward-btn:hover .label {
  color: #0078d4; /* Outlook's brand blue */
}

/* Red for destructive actions: Delete, Report, Junk */
.delete-btn:hover .material-symbols-outlined,
.delete-btn:hover .label,
.report-btn:hover .material-symbols-outlined,
.report-btn:hover .label,
.spam-btn:hover .material-symbols-outlined,
.spam-btn:hover .label {
  color: #d9534f;
}

/* Yellow for Starred */
.starred-btn:hover .fa-regular,
.starred-btn:hover .label {
  color: #f5c518;
}

/* Vertical separator line */
.separator {
  height: 32px;
  width: 1px;
  background-color: #e1e1e1;
  margin: 0 8px;
}

/* --- Delete Dialog (minor style tweaks) --- */
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
  justify-content: flex-end;
  gap: 10px;
}
.dialog-btn {
  padding: 8px 16px;
  border: 1px solid transparent;
  border-radius: 4px;
  cursor: pointer;
  font-weight: 600;
}
.confirm-btn {
  background: #d9534f;
  border-color: #d9534f;
  color: #fff;
}
.confirm-btn:hover {
  background: #c9302c;
  border-color: #c9302c;
}
.cancel-btn {
  background: #e2e8f0;
  color: #1a202c;
}
.cancel-btn:hover {
  background: #cbd5e0;
}
</style>
