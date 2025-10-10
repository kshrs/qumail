<template>
  <div class="compose-view">
    <div class="compose-header">
      <h3>New Message</h3>
    </div>
    <div class="compose-body">
      <div class="form-group">
        <label for="to">To</label>
        <input id="to" type="text" v-model="to" />
        <div class="cc-bcc">
          <span @click="showCc = !showCc" :class="{ active: showCc }">Cc</span>
          <span @click="showBcc = !showBcc" :class="{ active: showBcc }">Bcc</span>
        </div>
      </div>
      <div v-if="showCc" class="form-group">
        <label for="cc">Cc</label>
        <input id="cc" type="text" v-model="cc" />
      </div>
      <div v-if="showBcc" class="form-group">
        <label for="bcc">Bcc</label>
        <input id="bcc" type="text" v-model="bcc" />
      </div>
      <div class="form-group">
        <input id="subject" type="text" v-model="subject" placeholder="Subject" />
      </div>
      <textarea v-model="body" class="email-body" placeholder="Write something..."></textarea>
      <div v-if="attachments.length > 0" class="attachments-container">
        <div v-for="(file, index) in attachments" :key="index" class="attachment-pill">
          <i class="fa-solid fa-paperclip"></i>
          <span>{{ file.name }}</span>
          <button @click="removeAttachment(index)" class="remove-btn">&times;</button>
        </div>
      </div>
    </div>
    <input
      type="file"
      multiple
      ref="fileInput"
      @change="handleFileSelection"
      style="display: none"
    />
    <div class="compose-footer">
      <button class="send-btn" @click="sendEmail">
        <span>Send</span>
        <i class="fa-solid fa-paper-plane"></i>
      </button>
      <div class="footer-options">
        <button class="options-btn" @click="triggerFileInput" title="Attach files">
          <i class="fa-solid fa-paperclip"></i>
        </button>
        <button class="options-btn" @click="showConfirmDialog = true" title="Discard draft">
          <i class="fa-solid fa-trash"></i>
        </button>
      </div>
    </div>

    <Transition name="fade">
      <div v-if="showConfirmDialog" class="modal-overlay" @click.self="showConfirmDialog = false">
        <div class="confirm-dialog">
          <h3>Discard Draft?</h3>
          <p>Are you sure you want to discard this draft? This action cannot be undone.</p>
          <div class="dialog-actions">
            <button class="btn-cancel" @click="showConfirmDialog = false">Cancel</button>
            <button class="btn-confirm" @click="confirmDiscard">Discard</button>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { SendEmail } from '../../wailsjs/go/main/App';

const emit = defineEmits(['close']);

const to = ref('');
const cc = ref('');
const bcc = ref('');
const subject = ref('');
const body = ref('');
const showCc = ref(false);
const showBcc = ref(false);

const attachments = ref([]);
const fileInput = ref(null);

// CHANGE 3: New state variable to control the dialog's visibility
const showConfirmDialog = ref(false);

const triggerFileInput = () => {
  fileInput.value.click();
};

const handleFileSelection = (event) => {
  const selectedFiles = Array.from(event.target.files);
  attachments.value.push(...selectedFiles);
  event.target.value = '';
};

const removeAttachment = (index) => {
  attachments.value.splice(index, 1);
};

// CHANGE 4: New function to handle the confirmation action
const confirmDiscard = () => {
  showConfirmDialog.value = false; // Hide the dialog
  emit('close'); // Perform the original close action
};

const sendEmail = async () => {
  if (!to.value.trim()) {
    alert('Please enter at least one recipient in the "To" field.');
    return;
  }

  try {
    const result = await SendEmail(
    to.value,
    cc.value,
    bcc.value,
    subject.value,
    body.value
    );
  } catch (err) {
    alert(`Error sending Email: ${err}`);
    console.error(err)
  }
  const formData = new FormData();
  formData.append('to', to.value);
  formData.append('cc', cc.value);
  formData.append('bcc', bcc.value);
  formData.append('subject', subject.value);
  formData.append('body', body.value);
  attachments.value.forEach((file) => {
    formData.append('attachments', file);
  });
  console.log('✅ Email data is ready to be sent!');
  console.log('Here is the FormData object that would be sent to the backend:');
  for (let [key, value] of formData.entries()) {
    console.log(`${key}:`, value);
  }
  alert(
    'Email data has been prepared and logged to the console. Check the developer tools (F12) to see the FormData object.'
  );
  emit('close');
};
</script>

<style src ="../assets/Compose.css">
</style>
