<template>
  <div class="main-emaillist">

  <div class="inbox-container">

    <div class="inbox-toolbar">
      <div class="toolbar-left">

        <input
          type="checkbox"
          v-model="selectAll"
          class="email-checkbox"
          title="Select all"
        />


        <button
          @click="refreshEmails"
          class="options-btn"
          title="Refresh"
        >
          <i class="fa-solid fa-rotate-right"></i>
        </button>


        <div class="more-options-container">
          <button
            @click="isMoreMenuOpen = !isMoreMenuOpen"
            class="options-btn"
            title="More"
          >
            <i class="fa-solid fa-ellipsis-vertical"></i>
          </button>

        </div>
      </div>

      <div class="toolbar-right">
        <span>1-50 of {{ emails.length }}</span>
        <button class="options-btn" title="Newer">
          <i class="fa-solid fa-chevron-left"></i>
        </button>
        <button class="options-btn" title="Older">
          <i class="fa-solid fa-chevron-right"></i>
        </button>
      </div>
    </div>

    <div class="email-list" :class="{ loading: isLoading }">

      <div v-if="isLoading" class="loader">
        <i class="fa-solid fa-spinner fa-spin"></i>
        <span>Loading</span>
      </div>

      <div
        v-for="email in emails"
        :key="email.seqNum"
        class="email-row"
        @click="$emit('viewEmail', email)"
        :class="{
          'is-read': email.isRead,
          'is-selected': email.isSelected,
        }"
      >
        
        <div class="email-actions">
          <input
            type="checkbox"
            :checked="email.isSelected"
            @change="toggleSelect(email)"
            class="email-checkbox"
          />
          <button
            @click.stop="toggleStar(email)"
            class="star-btn"
            :title="email.isStarred ? 'Starred' : 'Not starred'"
          >
            <i
              :class="email.isStarred ? 'fa-solid fa-star' : 'fa-regular fa-star'"
            ></i>
          </button>
        </div>

       
        <div class="email-sender">
          <span>{{ email.from }}</span>
        </div>

      
        <div class="email-content">
          <span class="email-subject">{{ email.subject }}</span>
        </div>

       
        <div class="email-date">
          <span>{{ email.date }}</span>
        </div>
      </div>
    </div>
  </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue';
// 1. Import all the necessary backend functions
import { FetchEmails, ToggleStarred } from '../../wailsjs/go/main/App';
import { EventsOn } from '../../wailsjs/runtime/runtime';

const props = defineProps({
  section: {
    type: String,
    required: true,
  },
  isBackendReady: {
    type: Boolean,
    required: true,
  },
});

const emit = defineEmits(['viewEmail']);

const emails = ref([]);
const isLoading = ref(false);
const isMoreMenuOpen = ref(false);

const selectAll = computed({
  get: () => emails.value.length > 0 && emails.value.every(e => e.isSelected),
  set: (value) => emails.value.forEach(e => (e.isSelected = value)),
});

const loadEmails = async () => {
  isLoading.value = true;
  isMoreMenuOpen.value = false;
  try {
    const fetchedEmails = await FetchEmails(15, props.section);
    emails.value = (fetchedEmails || []).map(email => ({
      ...email,
      isSelected: false,
    })).reverse();
  } catch (err) {
    console.error(`Error fetching emails for section ${props.section}:`, err);
  } finally {
    isLoading.value = false;
  }
};

const refreshEmails = async () => {
  await loadEmails();
};

const toggleRead = async (email) => {
  try {
    // Pass the correct current read status
    await ToggleRead(email.seqNum, email.isRead);
    email.isRead = !email.isRead;
  } catch (err) {
    console.error("Error toggling read status:", err);
  }
};

const markAllAsRead = async () => {
  isMoreMenuOpen.value = false;
  try {
    await MarkAllAsRead();
    // Refresh the list to see the changes
    await loadEmails();
  } catch (err) {
    console.error("Error marking all as read:", err);
  }
};

// --- THIS IS THE CORRECTED FUNCTION ---
const toggleStar = async (email) => {
  try {
    await ToggleStarred(email.seqNum, email.isStarred);
    // 2. Only update the UI after the backend call is successful
    email.isStarred = !email.isStarred;
  } catch (err) {
    console.error("Error toggling star:", err);
    // Optional: revert the UI change if the backend fails
    // email.isStarred = !email.isStarred; 
  }
};

const toggleSelect = (email) => {
  email.isSelected = !email.isSelected;
};

// Load initial emails only after the backend is ready
onMounted(() => {
  // The component is re-mounted for every new section because of the :key.
  // Just load the emails directly.
  loadEmails();
});

// The watcher is still good to have, so you can leave it as is.
//watch(() => props.section, (newSection) => {
//  if (newSection) {
//    loadEmails();
//  }
//});

watch(() => props.isBackendReady, (isReady) => {
  if (isReady) {
    loadEmails();
  }
});

// Expose the loadEmails function for the parent to use
defineExpose({
  loadEmails,
});
</script>

<style scoped>
.main-emaillist {
    border-radius: 25px;
}
.inbox-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background-color: #121212; /* Zen: Main Background */
  font-family: 'Inter', sans-serif;
  color: #D1CFC0; /* Zen: Primary Text */
}

.inbox-toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 16px;
  background-color: #1E1E1E; /* Zen: UI Surface */
  border-bottom: 1px solid #2a2a2a; /* Zen: Subtle Border */
}

.more-menu {
  color: #D1CFC0;
  border-radius: 25px;
}

.more-menu:hover {
  background-color: #333333;
  color: #D1CFC0;
  border-radius: 25px;
  transition: background 0.28s;
}

.toolbar-left,
.toolbar-right {
  display: flex;
  align-items: center;
  gap: 10px;
  color: #D1CFC0;
}

.options-btn {
  border: none;
  background: transparent;
  cursor: pointer;
  font-size: 16px;
  padding: 6px;
  color: #D1CFC0;
}

.email-checkbox {
  cursor: pointer;
  accent-color: #F76F53; /* Zen: Melon Accent for checkbox */
  zoom: 1.4;
}

.email-list {
  flex: 1;
  overflow-y: auto;
  background: #121212; /* Zen: Main Background */
}

.email-row {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  border-bottom: 1px solid #2a2a2a; /* Zen: Subtle Border */
  cursor: pointer;
  transition: background-color 0.2s;
  opacity: 1;
}

.email-row:hover {
  background-color: rgba(247, 111, 83, 0.1); /* Zen: Melon Accent Hover */
}

.email-actions {
  display: flex;
  align-items: center;
  gap: 10px;
  width: 60px;
}

.email-sender {
  font-weight: 600;
  color: #D1CFC0; /* Zen: Primary Text */
  width: 180px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.email-content {
  flex: 1;
  display: flex;
  align-items: center;
  overflow: hidden;
}

.email-subject {
  font-weight: 500;
  color: #8A8A8A; /* Zen: Secondary Text */
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  padding: 0 8px;
}

.email-date {
  width: 100px;
  text-align: right;
  color: #8A8A8A; /* Zen: Secondary Text */
}

.star-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: #8A8A8A; /* Zen: Secondary Text */
  font-size: 16px;
}

.star-btn:hover {
  color: #FFB900; /* Vibrant yellow for hover */
}

.fa-solid.fa-star {
  color: #FFB900; /* Vibrant yellow for selected star */
}

.email-row.is-selected {
  background-color: rgba(247, 111, 83, 0.2); /* More prominent accent for selection */
}

.email-row.is-read {
  opacity: 0.6; /* Dim read emails */
}

.email-row.is-read .email-sender {
  font-weight: 400; /* Make sender less bold when read */
}

.loader {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 30px;
  color: #F76F53; /* Zen: Melon Accent for loader */
  font-weight: 500;
}

.loader i {
  margin-right: 10px;
}
</style>
