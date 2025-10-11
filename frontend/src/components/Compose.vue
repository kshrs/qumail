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
      <button class="send-btn" @click="handleSendClick">
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
    
    <Transition name="fade">
      <div v-if="showSendConfirmDialog" class="modal-overlay" @click.self="showSendConfirmDialog = false">
        <div class="confirm-dialog">
          <h3>Send Message?</h3>
          <p>Please confirm you want to send this email to the specified recipients.</p>
          <div class="dialog-actions">
            <button class="btn-cancel" @click="showSendConfirmDialog = false">Cancel</button>
            <button class="btn-confirm-send" @click="confirmSend">Send</button>
          </div>
        </div>
      </div>
    </Transition>
    <Transition name="fade">
  <div v-if="showAlert" class="modal-overlay" @click.self="showAlert = false">
    <div class="confirm-dialog">
      <h3>{{ alertTitle }}</h3>
      <p>{{ alertMessage }}</p>
      <div class="dialog-actions">
        <button class="btn-alert-ok" @click="showAlert = false">OK</button>
      </div>
    </div>
  </div>
</Transition>
  </div>
</template>


<script setup>
import { ref } from 'vue';

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


const showConfirmDialog = ref(false);

const showSendConfirmDialog = ref(false);
const showAlert = ref(false);
const alertTitle = ref('');
const alertMessage = ref('');

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
const handleSendClick = () => {
  if (!to.value.trim()) {
   
    alertTitle.value = 'Recipient Missing';
    alertMessage.value = 'Please add at least one recipient to the "To" field before sending.';
    showAlert.value = true;
  } else {
 
    showSendConfirmDialog.value = true;
  }
};

const confirmDiscard = () => {

  showConfirmDialog.value = false;
  emit('close');
};

const confirmSend = () => {
  showSendConfirmDialog.value = false;
  sendEmail();
};
const sendEmail = () => {

  const formData = new FormData();
  formData.append('to', to.value);
  formData.append('cc', cc.value);
  formData.append('bcc', bcc.value);
  formData.append('subject', subject.value);
  formData.append('body', body.value);
  attachments.value.forEach((file) => {
    formData.append('attachments', file);
  });
  console.log('Email data is ready to be sent!');
  console.log('Here is the FormData object that would be sent to the backend:');
  for (let [key, value] of formData.entries()) {
    console.log(`${key}:`, value);
  }
  emit('close');
};
</script>

<style src="../assets/styles/Compose.css"></style>