<template>
  <div class="compose-view">
    <div class="compose-header">
      <h3>New Message</h3>
    <div class="dropdown-container">
    
      <button class="security-level-btn" @click="toggleDropdown">
        <span>{{ selectedLevel.label }}</span>
        <i class="fa-solid fa-chevron-down"></i>
      </button>
    
      <Transition name="fade">
        <div v-if="dropdownOpen" class="dropdown-menu">
          
          <div 
            v-for="level in levels" 
            :key="level.value" 
            class="dropdown-item" 
            @click="selectLevel(level)"> <strong>{{ level.label }}</strong>
            <p>{{ level.description }}</p>
          </div>
    
        </div>
      </Transition>
    </div>
      <div class="compose-controls">
        <button class="close-btn" @click="$emit('close')" title="Close">×</button>
      </div>
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
    <div class="compose-footer">
      <button class="send-btn" @click="handleSendClick">
          <span>{{ sendingText }}</span>
        <i class="fa-solid fa-paper-plane"></i>
      </button>
      <div class="footer-options">
        <button class="options-btn" @click="addAttachments" title="Attach files">
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


<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue';
import { SendEmail, PickFiles } from '../../wailsjs/go/main/App';

const sendingText = ref('Send');

const emit = defineEmits(['close']);

const to = ref('');
const cc = ref('');
const bcc = ref('');
const subject = ref('');
const body = ref('');
const showCc = ref(false);
const showBcc = ref(false);

const attachments = ref<{ name: string, path: string }[]>([]);



const showConfirmDialog = ref(false);

const showSendConfirmDialog = ref(false);
const showAlert = ref(false);
const alertTitle = ref('');
const alertMessage = ref('');

// Security Level
// --- Dropdown state
const dropdownOpen = ref(false);          // whether dropdown is open or not
const levels = [{
  value: "OTP",
  label: "Level 1: One Time Pad",
  description: "Classical symmetric key encryption. Not Quantum Secure."
  },
  {
  value: "QKD_AES",
  label: "Level 2: QKD-Seeded AES",
  description: "Uses Quantum Key Distribution to seed the AES algorithm."
  },
  {
  value: "PQC",
  label: "Level 3: Post-Quantum Cryptography",
  description: "Uses algorithms resistant to qunatum attacks."
  },
  {
  value: "NONE",
  label: "Level 4: No Encryption",
  description: "Data is sent in plaintext without any security."
  }
];
const selectedLevel = ref(levels[3]);  // button label

function toggleDropdown(e: Event) {
  e.stopPropagation();
  dropdownOpen.value = !dropdownOpen.value;
};

function selectLevel(level: typeof levels[0]) {
  selectedLevel.value = level;
  dropdownOpen.value = false;
};

// --- close dropdown on outside click
function closeAll() { 
dropdownOpen.value = false;
};
onMounted(() => window.addEventListener('click', closeAll));
onBeforeUnmount(() => window.removeEventListener('click', closeAll));



// Attach files
const addAttachments = async () => {
  try {
    const filePaths = await PickFiles();
    
    if (filePaths) {
      // Create objects with a 'name' (for display) and 'path' (to send to Go)
      const newAttachments = filePaths.map(path => {
        const parts = path.split(/[\\/]/); // Handles both Windows and Unix paths
        return { name: parts[parts.length - 1], path: path };
      });
      attachments.value.push(...newAttachments);
    }
  } catch (err) {
    console.error("Error selecting files:", err);
  }
};

// Detach Files
const removeAttachment = (index: number) => {
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

// Send Email with attachments
const sendEmail = async () => {
  if (!to.value.trim()) {
    alert('Please enter at least one recipient in the "To" field.');
    return;
  }

  const attachmentPaths = attachments.value.map(file => file.path);
  try {
    sendingText.value = 'Sending ...'
    const result = await SendEmail(
    to.value,
    cc.value,
    bcc.value,
    subject.value,
    body.value,
    attachmentPaths,
    selectedLevel.value.value
    );
    emit('close');
  } catch (err) {
    alert(`Error sending Email: ${err}`);
    console.error(err)
  } finally {
    sendingText.value = 'Send';
  }
};
</script>

<style scoped>
.compose-view {
  width: 100%;
  height: 85%;
  background-color: #1E1E1E; /* Zen: UI Surface */
  color: #D1CFC0; /* Zen: Primary Text */
  border-radius: 16px;
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.3); /* Dark theme shadow */
  display: flex;
  flex-direction: column;
  overflow: hidden;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica,
    Arial, sans-serif;
}
.compose-controls {
  margin-left: 1vw;
  display: flex;
  gap: 8px;
  align-items: center;
}
.close-btn {
  background: transparent;
  border: none;
  cursor: pointer;
  font-size: 20px;
  padding: 6px 8px;
  border-radius: 6px;
  color: #D1CFC0;
}
.close-btn:hover { background: #333333; }
.compose-header {
  padding: 16px 24px;
  border-bottom: 1px solid #2a2a2a; /* Zen: Subtle Border */
  flex-shrink: 0;
  display: flex;
  align-items: center;
}
.compose-body {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  overflow-y: auto;
  padding: 8px 24px;
}

.form-group {
  display: flex;
  align-items: center;
  border-bottom: 1px solid #2a2a2a; /* Zen: Subtle Border */
  padding: 8px 0;
}
.form-group label {
  font-size: 14px;
  color: #8A8A8A; /* Zen: Secondary Text */
  margin-right: 16px;
  font-weight: 500;
}
.form-group input {
  flex-grow: 1;
  border: none;
  outline: none;
  background: transparent;
  font-size: 15px;
  padding: 8px 0;
  color: #D1CFC0; /* Zen: Primary Text */
}
.form-group input::placeholder {
  color: #555;
}
.cc-bcc span.active {
  color: #F79253; /* Zen: Melon Accent */
}
.cc-bcc {
  display: flex;
  gap: 16px;
}
.cc-bcc span {
  color: #8A8A8A; /* Zen: Secondary Text */
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  padding: 4px 0;
  transition: color 0.2s ease;
}
.cc-bcc span:hover {
  color: #F79253; /* Zen: Melon Accent */
}
.email-body {
  flex-grow: 1;
  border: none;
  outline: none;
  resize: none;
  padding: 20px 0;
  font-size: 16px;
  line-height: 1.7;
  color: #D1CFC0; /* Zen: Primary Text */
  font-family: inherit;
  background-color: transparent;
}
.email-body::placeholder {
  color: #555;
}
.attachments-container {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  padding: 16px 0;
  border-top: 1px solid #2a2a2a; /* Zen: Subtle Border */
  flex-shrink: 0;
}
.attachment-pill {
  display: flex;
  align-items: center;
  gap: 8px;
  background-color: #333333;
  border: 1px solid #4a4a4a;
  color: #D1CFC0;
  font-size: 13px;
  font-weight: 500;
  border-radius: 16px;
  padding: 6px 8px 6px 12px;
  transition: background-color 0.2s ease;
}
.attachment-pill:hover {
  background-color: #4a4a4a;
}
.attachment-pill .fa-paperclip {
  font-size: 12px;
}
.remove-btn {
  background: #555;
  border: none;
  color: #D1CFC0;
  cursor: pointer;
  font-size: 14px;
  font-weight: bold;
  width: 20px;
  height: 20px;
  line-height: 20px;
  text-align: center;
  border-radius: 50%;
  transition: all 0.2s ease;
}
.remove-btn:hover {
  background-color: #c0392b; /* Destructive Red */
  color: white;
  transform: rotate(90deg);
}
.compose-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 24px;
  border-top: 1px solid #2a2a2a; /* Zen: Subtle Border */
  background-color: #1E1E1E;
  flex-shrink: 0;
}
.send-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  background-color: #F79253; /* Zen: Melon Accent */
  color: #363636; /* Dark text for contrast */
  border: none;
  padding: 10px 20px;
  border-radius: 8px;
  cursor: pointer;
  font-size: 15px;
  font-weight: 600;
  transition: background-color 0.2s ease, transform 0.1s ease;
}
.send-btn:hover {
  background-color: #e65a40;
}
.send-btn:active {
  transform: scale(0.98);
}
.footer-options {
  display: flex;
  align-items: center;
  gap: 8px;
}
.options-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: #8A8A8A; /* Zen: Secondary Text */
  font-size: 18px;
  width: 36px;
  height: 36px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.2s ease, color 0.2s ease;
}
.options-btn:hover {
  background-color: #333333;
  color: #D1CFC0;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0,0,0, 0.6);
  backdrop-filter: blur(2px);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.confirm-dialog {
  background: #1E1E1E;
  border: 1px solid #2a2a2a;
  padding: 32px;
  border-radius: 16px;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.4);
  width: 100%;
  max-width: 400px;
  text-align: center;
  transform: scale(1);
}

.confirm-dialog h3 {
  margin-top: 0;
  margin-bottom: 12px;
  font-size: 20px;
  color: #D1CFC0;
}

.confirm-dialog p {
  margin-bottom: 24px;
  color: #8A8A8A;
  font-size: 15px;
  line-height: 1.6;
}

.dialog-actions {
  display: flex;
  justify-content: flex-end;
  gap: 16px;
}

.dialog-actions button {
  border: none;
  padding: 10px 24px;
  border-radius: 8px;
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
}

.btn-cancel {
  background-color: #333333;
  color: #D1CFC0;
}
.btn-cancel:hover {
  background-color: #4a4a4a;
}

.btn-confirm {
  background-color: #F79253; /* Zen: Melon Accent */
  color: #1E1E1E;
}
.btn-confirm:hover {
  background-color: #e65a40;
}

.btn-confirm-send {
  background-color: #111111;
  color: #4a4a4a;
}
.btn-confirm-send:hover {
  background-color: #d1cfc0;
  color: #1e1e1e;
}

/* Vue Transition Styles */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-active .confirm-dialog,
.fade-leave-active .confirm-dialog {
  transition: transform 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0.0;
}
.fade-enter-from .confirm-dialog,
.fade-leave-to .confirm-dialog {
  transform: scale(0.95);
}

/* --- Dropdown Container Theming --- */
.dropdown-container {
  position: relative;
  width: auto;
  margin-bottom: auto;
  margin-top: auto;
  margin-left: auto;
}

.security-level-btn {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  padding: 10px 12px;
  font-size: 1em;
  text-align: left;
  background-color: #363636;
  color: #D1CFC0;
  border: 1px solid #4a4a4a;
  border-radius: 6px;
  cursor: pointer;
  transition: background-color 0.2s ease;
  outline: none;
}

.security-level-btn:hover {
  background-color: #333333;
}

.security-level-btn .fa-chevron-down {
  transition: transform 0.2s ease;
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  left: 0;
  width: 100%;
  background-color: #111111;
  border: 1px solid #2a2a2a;
  border-radius: 6px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.3);
  margin-top: 5px;
  z-index: 1000;
  max-height: 300px;
  overflow-y: auto;
  text-align: left;
}

.dropdown-item {
  padding: 12px 15px;
  cursor: pointer;
  border-bottom: 1px solid #2a2a2a;
  transition: background-color 0.2s ease;
}

.dropdown-item:last-child {
  border-bottom: none;
}

.dropdown-item:hover {
    background-color: #1e1e1e;
}

.dropdown-item strong {
  display: block;
  color: #D1CFC0;
}

.dropdown-item p {
  display: none;
  font-size: 0.875em;
  color: #8A8A8A;
  margin: 5px 0 0 0;
}

.dropdown-item:hover p {
  display: block;
}

/* Unused classes from the original file have been omitted for brevity */
</style>
