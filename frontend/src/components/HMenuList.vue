<template>
  <ul class="List">
    <li @click="$emit('open-inbox')">
      <a href="#">
        <span class="material-symbols-outlined">inbox</span>
        Inbox
      </a>
    </li>
    <li>
      <a href="#">
        <span class="material-symbols-outlined">send</span>
        Sent Items
      </a>
    </li>
    <li>
      <a href="#">
        <span class="material-symbols-outlined">draft</span>
        Drafts
      </a>
    </li>
  </ul>
</template>

<script setup>
defineEmits(['open-inbox']);
</script>

<style>
* {
  padding: 0px;
  margin: 0px;
}
.List {
  list-style: none;
  padding: 0;
  margin-top: 8px;
}
.List li {
  margin-bottom: 20px;
  padding: 5px;
  border-radius: 3px; /* Add radius here */
  cursor: pointer;   /* Make it look clickable */
}
.List li a {
  color: black;
  text-decoration: none;
  font-size: 20px;
  display: flex;
}

/* FIX: Target the 'li' element for hover, active, and focus, 
  not the whole list.
*/
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
  font-size: 22px;
}
</style>
