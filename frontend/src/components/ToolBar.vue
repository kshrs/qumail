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
function closeAll() { dropdownOpen.value = false}
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
        <button class="feature-btn report-btn" @click="goTo('#')" title="Report">
          
        <span class="material-symbols-outlined" id="report" style="color: darkred;">report
        </span>
        </button>
        <button class="feature-btn reply-btn" @click="goTo('#')" title="Reply">
        
        <span class="material-symbols-outlined " style="color: orange;">
reply
</span>
        </button>
        <button class="feature-btn reply-all-btn" @click="goTo('#')" title="Reply All">
          <span class="material-symbols-outlined" style="color: orange;">
reply_all
</span>
        </button>
        <button class="feature-btn forward-btn" @click="goTo('#')" title="Forward">
          <span class="material-symbols-outlined" style="color: green;">
forward
</span>
        </button>
        <button class="feature-btn mark-read-btn" @click="goTo('#')" title="Mark as Read">
          <span class="material-symbols-outlined" style="color: brown;">
mark_as_unread
</span>
        </button>
        <button class="feature-btn starred-btn" @click="goTo('#')" title="Starred">
         <i class="fa-regular fa-star" style="color:#f5c518"></i>
        </button>
        <button class="feature-btn important-btn" @click="goTo('#')" title="Important">
          <span class="material-symbols-outlined" style="color: purple;">
exclamation
</span>
        </button>
        <button class="feature-btn spam-btn" @click="goTo('#')" title="Spam">
          <span class="material-symbols-outlined" style="color: red;">
dangerous
</span>
        </button>
        <button class="feature-btn delete-btn" @click="confirmDelete" title="Delete">
          <span class="material-symbols-outlined">
delete
</span>
        </button>
        <button class="feature-btn settings-btn" @click="goTo('#')" title="Settings">  
          <span class="material-symbols-outlined" style="color: blue;">
settings
</span>
        </button>
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

<style scoped>
.feature-buttons button {
  background: none;       /* removes button background */
  border: none;           /* removes border */
  padding: 0;             /* removes default padding */
  margin: 0;              /* removes default margin */
  color: black;         /* uses parent text color */
  font: inherit;          /* uses parent font */
  cursor: pointer;        /* keeps the pointer cursor */
}

.feature-buttons button:hover {
  text-decoration: underline;  /* optional hover effect */
}
.top-bar{
  background-color:#fafafa;
  height: 50px;
  padding-right: 16px;
  display:flex;
  align-items:center;
  justify-content:space-between;
}


</style>
