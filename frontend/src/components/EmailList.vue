<template>

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

          <Transition name="fade">
            <div v-if="isMoreMenuOpen" class="more-menu">
              <div @click="markAllAsRead" class="menu-item">
                Mark all as read
              </div>
            </div>
          </Transition>
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
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { FetchEmails, ToggleRead, ToggleStarred } from '../../wailsjs/go/main/App';

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
    const fetchedEmails = await FetchEmails(10);
    emails.value = fetchedEmails.map(email => ({
        ...email,
        isSelected: false,

    })).reverse();
  } catch (err) {
    console.error(err);
  } finally {
    isLoading.value = false;
  }
}

const refreshEmails = async () => {
    await loadEmails();
};

// Working on it
const toggleRead = async (email) => {
    await ToggleRead(email.seqNum, email.isStarred);
}

// Need to be worked on 
const markAllAsRead = () => {
  emails.value.forEach(e => (e.isRead = true));
  isMoreMenuOpen.value = false;
};

// Need to be worked on
const toggleStar = async (email) => {
  email.isStarred = !email.isStarred;
  await ToggleStarred(email.seqNum, !email.isStarred);
};

const toggleSelect = (email) => {
  email.isSelected = !email.isSelected;
};

onMounted(() => {
    loadEmails();
});
</script>

<style src="../assets/styles/EmailList.css"></style>
