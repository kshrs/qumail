<template>
  <div class="inbox-container">
    <div class="inbox-toolbar">
      <div class="toolbar-left">
        <input type="checkbox" v-model="selectAll" class="email-checkbox" title="Select all" />
        <button @click="refreshEmails" class="options-btn" title="Refresh"><i class="fa-solid fa-rotate-right"></i></button>
        
        <div class="more-options-container">
          <button @click="isMoreMenuOpen = !isMoreMenuOpen" class="options-btn" title="More"><i class="fa-solid fa-ellipsis-vertical"></i></button>
          
          <Transition name="fade">
            <div v-if="isMoreMenuOpen" class="more-menu">
              <div @click="markAllAsRead" class="menu-item">Mark all as read</div>
            </div>
          </Transition>
        </div>
      </div>

      <div class="toolbar-right">
        <span>1-50 of 4,449</span>
        <button class="options-btn" title="Newer"><i class="fa-solid fa-chevron-left"></i></button>
        <button class="options-btn" title="Older"><i class="fa-solid fa-chevron-right"></i></button>
      </div>
    </div>

    <div class="email-list" :class="{ loading: isLoading }">
      <div v-if="isLoading" class="loader">
        <i class="fa-solid fa-spinner fa-spin"></i>
        <span>Loading</span>
      </div>
      <EmailListItem
        v-for="email in emails"
        :key="email.id"
        :email="email"
        @toggle-star="handleToggleStar"
        @toggle-select="handleToggleSelect"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import EmailListItem from './EmailList.vue';


const isLoading = ref(false);
const isMoreMenuOpen = ref(false);
const emails = ref([
  { id: 1, sender: 'Medium Daily Digest', subject: "Elon Musk Doesn't Understand AI", snippet: 'I Will Lockett | G Ashwin Balaji Stories', timestamp: '07:00', isRead: false, isStarred: false, isSelected: false },
  { id: 2, sender: 'LinkedIn', subject: 'Ashwin, thanks for being a valued member', snippet: 'This is the preview text visible in notifications', timestamp: '03:40', isRead: false, isStarred: true, isSelected: false },
  { id: 3, sender: 'Gamma', subject: 'Pay with UPI now available', snippet: "We're excited to share that Unified Payments Interface (UPI) is now a payment option", timestamp: '01:35', isRead: false, isStarred: false, isSelected: false },
  { id: 4, sender: 'Coursera', subject: 'Jumpstart your JavaScript journey with Microsoft', snippet: 'Master front-end development fundamentals by building web applications', timestamp: '9 Oct', isRead: true, isStarred: false, isSelected: false },
  { id: 5, sender: 'Google', subject: 'Security alert for gashwinbalaji10@gmail.com', snippet: 'This is a copy of a security alert sent to gashwinbalaji10@gmail.com', timestamp: '9 Oct', isRead: true, isStarred: false, isSelected: false },
  { id: 6, sender: 'Kishor 2', subject: 'Re: [kshrs/qumail] Added compose part', snippet: 'Compose.vue, Compose.css) (PR #1) · Merged #1 into main. – Reply to this email directly', timestamp: '9 Oct', isRead: true, isStarred: false, isSelected: false },
  { id: 7, sender: 'Adobe for Photographers', subject: 'Learn the trick behind main character magic', snippet: 'Deculter your masterpiece with Generative Remove and Quick Actions', timestamp: '9 Oct', isRead: true, isStarred: true, isSelected: false },
  { id: 8, sender: 'Tata Consultancy Services', subject: 'Tata Consultancy Services is hiring', snippet: 'See new openings', timestamp: '8 Oct', isRead: true, isStarred: false, isSelected: false },
]);

const selectAll = computed({
 
  get: () => emails.value.length > 0 && emails.value.every(e => e.isSelected),

  set: (value) => {
    emails.value.forEach(email => email.isSelected = value);
  }
});


const refreshEmails = () => {
  isLoading.value = true;
  isMoreMenuOpen.value = false; 
  setTimeout(() => {

    const firstUnread = emails.value.find(e => e.isRead);
    if (firstUnread) firstUnread.isRead = false;
    
    emails.value.unshift({ id: Date.now(), sender: 'Kishor S', subject: 'Jolly Work!', snippet: 'I had completed my work!', timestamp: 'Now', isRead: false, isStarred: false, isSelected: false });

    isLoading.value = false;
  }, 1500);
};

const markAllAsRead = () => {
  emails.value.forEach(email => email.isRead = true);
  isMoreMenuOpen.value = false; 
};

const handleToggleStar = (emailId) => {
  const email = emails.value.find(e => e.id === emailId);
  if (email) email.isStarred = !email.isStarred;
};

const handleToggleSelect = ({ id, selected }) => {
  const email = emails.value.find(e => e.id === id);
  if (email) email.isSelected = selected;
};
</script>

<style src = "../assets/EmailList.css"> </style>