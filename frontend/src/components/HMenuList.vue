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
  padding: 5px;
  border-radius: 5px; /* Add radius here */
  cursor: pointer;   /* Make it look clickable */
  text-align: left;
  padding: 10px;
  font-size: 20px;
  margin-top: 1px;
}


.List li:hover {
  background: #e3f0ff;
}

.List li:active {
  background: #ccdff6;
}

.List li:focus-within { /* Use focus-within for better accessibility */
  background: #ccdff6;
  box-shadow: 0 0 0 4px #ccdff6;
}


.material-symbols-outlined {
  margin-right: 6px;
  margin-top: 1px;
  font-variation-settings: 'FILL' 1, 'wght' 300, 'GRAD' 0, 'opsz' 48;
  font-size: 20px;

}

/* 3. Add a style for the new .active class */
.List li.active {
  background: #ccdff6;
  font-weight: 600; /* Make the text bold */
}
</style>
