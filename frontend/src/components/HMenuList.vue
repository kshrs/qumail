<template>
  <ul class="List">
    <li :class="{ active: activeView === 'inbox' }" @click="$emit('open-inbox')">
      <span class="material-symbols-outlined">inbox</span>
      Inbox
    </li>
    
    <li :class="{ active: activeView === 'sent' }" @click="$emit('open-sent')">
      <span class="material-symbols-outlined">send</span>
      Sent
    </li>
    
    <li :class="{ active: activeView === 'spam' }" @click="$emit('open-spam')">
      <span class="material-symbols-outlined">dangerous</span>
      Spam
    </li>
    <li :class="{ active: activeView === 'drafts' }" @click="$emit('open-drafts')">
      <span class="material-symbols-outlined">draft</span>
      Drafts
    </li>
    <li :class="{ active: activeView === 'trash' }" @click="$emit('open-trash')">
      <span class="material-symbols-outlined">delete</span>
      Trash
    </li>
    <li :class="{ active: activeView === 'allmail' }" @click="$emit('open-allmail')">
      <span class="material-symbols-outlined">all_inbox</span>
      All Mail
    </li>
    <li :class="{ active: activeView === 'starred' }" @click="$emit('open-starred')">
      <span class="material-symbols-outlined">star</span>
      Starred
    </li>
    <li :class="{ active: activeView === 'important' }" @click="$emit('open-important')">
      <span class="material-symbols-outlined">flag_2</span>
      Important
    </li>
  </ul>
  <div class="google-user-profile">
    <span class="material-symbols-outlined">
    account_circle
    </span>
    <span class="material-symbols-outlined">
    <label> ProjectBase</label>
    </span>
  </div>
</template>

<script setup>
// 1. Define the prop the component will receive
defineProps({
  activeView: String
});

// 2. Define all possible events this component can emit
defineEmits(['open-inbox', 'open-sent', 'open-drafts', 'open-spam', 'open-trash', 'open-allmail', 'open-starred', 'open-important']);
</script>

<style>

* {
  padding: 0px;
  margin: 0px;
}

.List {
  -webkit-user-select: none;
  list-style: none;
  padding: 0;
  margin-top: 8px;
}

.List li {
  display: flex;
  align-items: center;
  padding: 10px;
  border-radius: 5px;
  cursor: pointer;
  text-align: left;
  font-size: 16px; /* Adjusted for better fit */
  margin-top: 1px;
  color: #D1CFC0; /* Zen Primary Text */
  transition: background-color 0.2s ease, color 0.2s ease;
}


.List li:hover {
  background: #363636; /* Slightly more prominent accent */
}

/* Style for the active state (when clicked or selected) */
.List li.active,
.List li:active {
  background: #363636; /* Slightly more prominent accent */
  color: #d1cfc0; /* Zen Accent Color */
  font-weight: 600; /* Make the text bold */
  border-left: 4px solid #f79253;
}

.List li:focus-within {
  box-shadow: 0 0 0 2px rgba(247, 111, 83, 0.4); /* Accessibility focus ring */
  outline: none;
}


.material-symbols-outlined {
  margin-right: 16px; /* Increased margin for better spacing */
  font-variation-settings: 'FILL' 0, 'wght' 300, 'GRAD' 0, 'opsz' 24;
  font-size: 22px;
}

/* Style for icons within an active list item */
.List li.active .material-symbols-outlined {
  font-variation-settings: 'FILL' 1, 'wght' 400, 'GRAD' 0, 'opsz' 24;
}

.google-user-profile {
  display: flex;
  align-items: center;
  margin-top: auto;
  margin-bottom: 10px;
  padding: 10px;
  color: #D1CFC0; /* Zen Primary Text */
}
.google-user-profile .material-symbols-outlined {
  font-size: 24px;
}
.google-user-profile .material-symbols-outlined label {
  font-size: 18px;
}
</style>
