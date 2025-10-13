<script lang="ts" setup>
import { ref } from 'vue'; // Removed unused imports for cleaner code

// --- New Mail Button
function newMail() {
  console.log("New Mail Clicked");
}


// --- Navigation helper
function goTo(href: string) {
  window.location.href = href;
}


// --- Delete button confirmation dialog
const showDeleteDialog = ref(false);
function confirmDelete() {
  showDeleteDialog.value = true;
}


function cancelDelete() {
  showDeleteDialog.value = false;
}

function deleteAction() {
  console.log("Delete action confirmed");
  showDeleteDialog.value = false;
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


<style>

/* Main container for the top bar */
.top-bar-container {
  /* The -webkit-user-select line has been removed */
  user-select: none;
}


/* The bar itself */
.top-bar {
  border: 2px solid #1e1e1e;
  border-radius: 10px;
  display: flex;
  align-items: center;
  padding: 4px;
  background-color: #1e1e1e; /* Dark background for the bar */
  margin-bottom: 10px;
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
  background-color: #333333; /* Lighter grey for hover feedback */
}


/* --- Icon and Label Base Styles --- */
.feature-btn .material-symbols-outlined,
.feature-btn .fa-regular {
  font-size: 22px;
  color: #d1d1d1; /* Light grey for icons */
  transition: color 0.2s ease-in-out;
}


.feature-btn .label {
  font-size: 11px;
  margin-top: 2px;
  color: #d1d1d1; /* Light grey for text */
  transition: color 0.2s ease-in-out;
}


/* --- Specific Hover Colors --- */


/* Bright Blue for primary actions: Reply, Reply All, Forward */
.reply-btn:hover .material-symbols-outlined,
.reply-btn:hover .label,
.reply-all-btn:hover .material-symbols-outlined,
.reply-all-btn:hover .label,
.forward-btn:hover .material-symbols-outlined,
.forward-btn:hover .label {
  color: #00aaff;
}


/* Vibrant Red for destructive actions: Delete, Report, Junk */
.delete-btn:hover .material-symbols-outlined,
.delete-btn:hover .label,
.report-btn:hover .material-symbols-outlined,
.report-btn:hover .label,
.spam-btn:hover .material-symbols-outlined,
.spam-btn:hover .label {
  color: #ff4d4d;
}


/* Rich Gold for Starred */
.starred-btn:hover .fa-regular,
.starred-btn:hover .label {
  color: #ffd700;
}


/* Vertical separator line */
.separator {
  height: 32px;
  width: 1px;
  background-color: #3a3a3a; /* Darker separator line */
  margin: 0 8px;
}


/* --- Delete Dialog (minor style tweaks) --- */
.delete-dialog {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: #2a2a2a; /* Dark dialog background */
  border: 1px solid #4a4a4a;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.5);
  z-index: 1000;
}

.delete-dialog-text {
  margin-bottom: 20px;
  font-size: 16px;
  color: #d1d1d1; /* Light text */
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
  background: #ff4d4d;
  border-color: #ff4d4d;
  color: #ffffff;
}

.confirm-btn:hover {
  background: #e60000;
  border-color: #e60000;
}

.cancel-btn {
  background: #4a4a4a;
  color: #d1d1d1;
}

.cancel-btn:hover {
  background: #5a5a5a;
}
</style>
