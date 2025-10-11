<script lang="ts" setup>
import { ref, reactive, onMounted, onBeforeUnmount } from 'vue'
import '../assets/styles/ToolBar.css'

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
