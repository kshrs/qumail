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

// --- Dropdown state
const dropdownOpen = ref(false)           // whether dropdown is open or not
const selectedLevel = ref("Select Security Level")  // button label
const levels = ["Level 1- Quantum Secure : use One Time Pad No Quantum security.",
                "Level 2- Quantum-aided AES: use Quantum keys as seed for AES.",
                "Level 3- Any other encryption (like PQC) may be given as option.",
                "Level 4- No Quantum security."]

function toggleDropdown(e: Event) {
  e.stopPropagation()
  dropdownOpen.value = !dropdownOpen.value
}

function selectLevel(level: string) {
  selectedLevel.value = level
  dropdownOpen.value = false
}

// --- close dropdown on outside click
function closeAll() { dropdownOpen.value = false }
onMounted(() => window.addEventListener('click', closeAll))
onBeforeUnmount(() => window.removeEventListener('click', closeAll))

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

      <!-- Dropdown Button -->
      <div class="dropdown-container">
        <button class="readable-select" @click.stop="toggleDropdown">{{ selectedLevel }}</button>

        <!-- Dropdown List -->
        <div class="select-dropdown" v-if="dropdownOpen">
          <div
            v-for="level in levels"
            :key="level"
            class="select-item"
            @click="selectLevel(level)"
          >
            {{ level }}
          </div>
        </div>
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
