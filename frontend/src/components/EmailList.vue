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

.inbox-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background-color: #fff;
  font-family: 'Inter', sans-serif;
}


.inbox-toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 16px;
  background-color: #f8f9fb;
  border-bottom: 1px solid #e5e7eb;
}

.more-menu {
  color: black;
  border-radius: 25px;
}

.more-menu:hover {
  background-color: #f6f6f6;
  color: black;
  border-radius: 25px;
  transition: background 0.28s;

}

.toolbar-left,
.toolbar-right {
  display: flex;
  align-items: center;
  gap: 10px;
  color: black;
}

.options-btn {
  border: none;
  background: transparent;
  cursor: pointer;
  font-size: 16px;
  padding: 6px;
}

.email-checkbox {
  cursor: pointer;
  zoom: 1.4;
}

.email-list {
  flex: 1;
  overflow-y: auto;
  background: #fff;
}

.email-row {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  border-bottom: 1px solid #eee;
  cursor: pointer;
  transition: background-color 0.2s;
  opacity: 1;
}

.email-row:hover {
  background-color: #f7f9fc;
}

.email-actions {
  display: flex;
  align-items: center;
  gap: 10px;
  width: 60px;
}

.email-sender {
  font-weight: 600;
  color: #333;
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
  color: #000;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.email-date {
  width: 100px;
  text-align: right;
  color: black;
}


.star-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: #999;
  font-size: 16px;
}


.star-btn:hover {
  color: #f5c518;
}

.fa-solid.fa-star {
  color: #f5c518;
}


.email-row.is-selected {
  background-color: #e8f0fe;
}


.email-row.is-read {
  opacity: 0.7;
}



.loader {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 30px;
  color: #666;
  font-weight: 500;
}

.loader i {
  margin-right: 10px;
}
</style>
