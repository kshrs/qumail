<template>
    <div class="email-view-container">
    <div class="email-view-header">
      <button @click="$emit('close')" class="back-btn" title="Back to Inbox"> <i class="fa-solid fa-arrow-left"></i> </button> <div class="header-actions">
        <button class="options-btn" title="Archive"><i class="fa-solid fa-box-archive"></i></button>
        <button class="options-btn" title="Delete" @click="handleDelete(props.email)"><i class="fa-solid fa-trash"></i></button>
        <button class="options-btn" title="More options"><i class="fa-solid fa-ellipsis-v"></i></button>
        <div class="decrypt-container">
        <button v-if="props.email.isEncrypted" class="decrypt-btn" @click="startOtpVerification">
            <span>Decrypt</span>
            <div class="material-symbols-outlined">
                lock_open_right
            </div>
        </button>
        </div>
      </div>
    </div>

    <div class="compose-body">
      <div class="form-group">
          <strong><label>From</label></strong>
        <div class="value-display">
          <strong>{{ props.email.from }}</strong>
        </div>
      </div>
      
      <div v-if="props.email.cc" class="form-group">
          <strong><label>Cc</label></strong>
        <div class="value-display">{{ props.email.cc }}</div>
      </div>

      <div class="form-group">
          <strong><label>Subject</label></strong>
        <div class="value-display subject-value">
          <strong>{{ props.email.subject }}</strong>
        </div>
      </div>
      
      <div class="email-body-content" v-html="props.email.body"></div>
      
      <div v-if="props.email.attachments && props.email.attachments.length > 0" class="attachments-container">
        <div v-for="(file, index) in props.email.attachments" :key="index" class="attachment-pill">
          <div class="material-symbols-outlined">attachment</div>
          <span>{{ file.name }} ({{ file.formattedSize }})</span>
          <button class="download-btn" title="Download" @click="downloadAttachment(file)">
              <div class="material-symbols-outlined">download</div>
          </button>
        </div>
      </div>
      <div class="email-footer">
      </div>
    <Transition name="fade">
      <div v-if="showOtpDialog" class="modal-overlay" @click.self="showOtpDialog = false">
        <div class="confirm-dialog">
          
          <div v-if="otpStep === 'enterPhone'">
            <h3>Two-Factor Authentication</h3>
            <p>For security, please enter your phone number to receive a verification code.</p>
            <input 
              type="tel" 
              v-model="phoneNumber" 
              class="key-input" 
              placeholder="+1234567890 (E.164 format)"
            />
            <div class="dialog-actions">
              <button class="btn-cancel" @click="showOtpDialog = false">Cancel</button>
              <button class="btn-confirm" @click="handleSendCode" :disabled="isLoading">
                {{ isLoading ? 'Sending...' : 'Send Code' }}
              </button>
            </div>
          </div>

          <div v-if="otpStep === 'enterCode'">
            <h3>Check Your Phone</h3>
            <p>We've sent a 5-digit code to <strong>{{ phoneNumber }}</strong>. Please enter it below.</p>
            <input 
              type="text" 
              v-model="otpCode" 
              class="key-input" 
              placeholder="12345"
              maxlength="6"
              @keyup.enter="handleVerifyCode"
            />
            <div class="dialog-actions">
              <button class="btn-cancel" @click="otpStep = 'enterPhone'">Back</button>
              <button class="btn-confirm" @click="handleVerifyCode" :disabled="isLoading">
                {{ isLoading ? 'Verifying...' : 'Verify & Decrypt' }}
              </button>
            </div>
          </div>

        </div>
      </div>
    </Transition>
    <Transition name="fade">
      <div v-if="showDecryptDialog" class="modal-overlay" @click.self="showDecryptDialog = false">
        <div class="confirm-dialog">
          <h3>Enter Decryption Key</h3>
          <p>Please enter the key required to decrypt this message content.</p>
          
          <input 
            type="text" 
            v-model="decryptionKey" 
            class="key-input" 
            placeholder="Enter your key..."
            @keyup.enter="handleDecryptConfirm(props.email)"
          />

          <div class="dialog-actions">
            <button class="btn-cancel" @click="showDecryptDialog = false">Cancel</button>
            <button class="btn-confirm" @click="handleDecryptConfirm(props.email)">Decrypt</button>
          </div>
        </div>
      </div>
    </Transition>   </div>
<Transition name="fade">
  <div v-if="showDeleteDialog" class="modal-overlay" @click.self="showDeleteDialog = false">
    <div class="confirm-dialog">
      <h3>Move to Trash?</h3>
      <p>Are you sure you want to move this conversation to the trash?</p>
      <div class="dialog-actions">
        <button class="btn-cancel" @click="showDeleteDialog = false">Cancel</button>
        <button class="btn-confirm btn-delete" @click="confirmDeleteAction">Move to Trash</button>
      </div>
    </div>
  </div>
</Transition>

  </div>
</template>

<script setup>
import { ref } from 'vue';
import { DownloadAttachment, Decrypt, DeleteSingleEmail, SendVerificationCode, CheckVerificationCode } from '../../wailsjs/go/main/App';
const showDecryptDialog = ref(false);
const decryptionKey = ref('');
const showDeleteDialog = ref(false);
const secret_key = ref('1234');


const showOtpDialog = ref(false);
const otpStep = ref('enterPhone'); // Controls which part of the dialog is visible
const phoneNumber = ref('');
const otpCode = ref('');
const isLoading = ref(false);

const tempSubject = ref('');
const tempBody = ref('');

// 1. Accept the new 'section' prop
const props = defineProps({
    email: {
        type: Object,
        required: true,
    },
    section: {
        type: String,
        required: true,
    }
});

const startOtpVerification = () => {
  // Reset state before showing
  otpStep.value = 'enterPhone';
  phoneNumber.value = '';
  otpCode.value = '';
  showOtpDialog.value = true;
};

const handleSendCode = async () => {
  if (!phoneNumber.value.trim()) {
    alert('Please enter a valid phone number.');
    return;
  }
  isLoading.value = true;
  try {
    console.log(phoneNumber.value);
    const result = await SendVerificationCode(phoneNumber.value);
    console.log(result); // "Verification code sent successfully."
    otpStep.value = 'enterCode'; // Move to the next step
  } catch (err) {
    alert(`Failed to send code: ${err}`);
  } finally {
    isLoading.value = false;
  }
};

const handleVerifyCode = async () => {
  if (!otpCode.value.trim()) {
    alert('Please enter the verification code.');
    return;
  }
  isLoading.value = true;
  try {
    const isVerified = await CheckVerificationCode(phoneNumber.value, otpCode.value);
    
    if (isVerified) {
      showOtpDialog.value = false;
      // Now you can call your original decryption logic
      // For simplicity, let's assume the key is handled by the backend now
      openDecryptDialog(); 
    } else {
      alert('Verification failed. The code is incorrect.');
    }
  } catch (err) {
    alert(`Verification error: ${err}`);
    showOtpDialog.value = false;
  } finally {
    isLoading.value = false;
  }
};


function handleDelete() {
  showDeleteDialog.value = true;
}

async function confirmDeleteAction() {
  try {
    const result = await DeleteSingleEmail(props.email.seqNum, props.section);
    console.log(result); // Log success
    showDeleteDialog.value = false; // Close the dialog
    emit('close'); // Go back to the inbox
  } catch (err) {
    alert(`Error deleting email: ${err}`); // Use a simple alert for errors
    showDeleteDialog.value = false;
  }
}

const emit = defineEmits(['close']);


const openDecryptDialog = () => {
  decryptionKey.value = ''; // Clear previous key
  showDecryptDialog.value = true;
};

// Step 2: This new function is called by the dialog's "Decrypt" button
const handleDecryptConfirm = async (email) => {
  if (!decryptionKey.value.trim()) {
    alert("Please enter a decryption key.");
    return;
  }

  if (secret_key.value != decryptionKey.value.trim()) {
    alert("Decryption Secret Key is wrong.");
    showDecryptDialog.value = false; // Close the dialog on success
    return;
  };


  //const tempSubject = props.email.subject.trim().replace("Encrypted:", "");
  //const tempBody = props.email.body.trim();
  
  try {
    decryptAction(email);
    showDecryptDialog.value = false; // Close the dialog on success
  } catch (err) {
    console.error(err);
    alert(`Decryption failed: ${err}`); // Show error to the user
    showDecryptDialog.value = false; // Also close dialog on failure
  }
};

const decryptAction = async (email) => {
    tempSubject.value = email.subject.trim().replace("Encrypted:", "");
    tempBody.value = email.body.trim();
    try {
       const decryptedBody = await Decrypt(tempBody.value, tempSubject.value); 
       console.log("Decrypted Message");
       console.log(decryptedBody);
       props.email.body = decryptedBody;
    } catch (err) {
        console.log(err);
    }
};

// 2. Update the download function to pass the section
const downloadAttachment = async (attachment) => {
    console.log("Trying to download");
    if (!props.email || !props.email.seqNum) return;
    console.log("Still Trying");

    try {
        console.log(`Downloading ${attachment.name} from section ${props.section}...`);
        
        // 3. Pass props.section as the third argument
        const result = await DownloadAttachment(props.email.seqNum, attachment.name, props.section);

        console.log(result);
    } catch (err) {
        console.error("Download failed:", err);
    }
}
</script>

<style scoped>
/* Base Styles */
.email-view-container {
  display: flex;
  flex-direction: column;
  height: calc(100vh - 9vh);
  background-color: #1E1E1E; /* Zen: UI Surface */
  color: #D1CFC0; /* Zen: Primary Text */
  font-family: 'Segoe UI', sans-serif;
  border: 2px solid #1e1e1e; /* Zen: Subtle Border */
  border-radius: 10px;
}

/* Header */
.email-view-header {
  display: flex;
  align-items: center;
  padding: 8px 16px;
  border-bottom: 1px solid #2a2a2a; /* Zen: Subtle Border */
  gap: 8px;
}

.header-actions {
  display: flex;
  width: 100%;
  gap: 8px;
}
.back-btn, .options-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: #D1CFC0; /* Zen: Primary Text */
  font-size: 18px;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.2s ease;
}
.back-btn:hover, .options-btn:hover { background-color: #333333; } /* Neutral Hover */

/* Body */
.compose-body {
  flex-grow: 1;
  display:flex;
  flex-direction: column;
  overflow-y: auto;
  padding: 8px 24px;
}

/* Form Group - Now a stacked, single-column layout */
.form-group {
  display: flex;
  flex-direction: row;
  align-items: flex-start; /* Align to top */
  border-bottom: 1px solid #2a2a2a; /* Zen: Subtle Border */
  padding: 12px 0;
}

/* The label is now styled as a small heading for the field below it */
.form-group label {
  font-size: 13px;
  color: #8A8A8A; /* Zen: Secondary Text */
  font-weight: 600;
  margin-bottom: 4px;
  width: 80px; /* Give it a fixed width */
}

/* This is the styled text that displays the value */
.value-display {
  flex-grow: 1;
  font-size: 15px;
  padding: 0;
  color: #D1CFC0; /* Zen: Primary Text */
  text-align: left;
  padding-left: 10px;
  margin-bottom: auto;
}

.subject-value {
  font-size: 16px;
  font-weight: bold;
}

/* Main Email Content */
.email-body-content {
  flex-grow: 1;
  margin-top: 15px;
  padding: 20px 15px;
  font-size: 16px;
  min-height: 500px;
  max-height: 700px;
  line-height: 1.7;
  width: 100%;
  color: #D1CFC0; /* Zen: Primary Text */
  overflow-wrap: break-word;
  text-align: left;
  background-color: #111111;
  border-radius: 15px;
}

.email-footer {
    background: #1e1e1e1;
    margin-bottom: 15px;
}

/* Attachments */
.attachments-container {
  display: flex; flex-wrap: wrap; gap: 10px;
  padding: 16px 0; border-top: 1px solid #2a2a2a;
}
.attachment-pill {
  display: flex; align-items: center; gap: 8px;
  background-color: #333333; border: 1px solid #4a4a4a;
  color: #D1CFC0; font-size: 13px; font-weight: 500;
  border-radius: 16px; padding: 6px 12px;
}
.download-btn {
  background: none; border: none; cursor: pointer; color: #D1CFC0;
  transition: color 0.2s;
}
.download-btn:hover { color: #F79253; } /* Zen: Melon Accent */

/* Footer */
.email-view-footer {
  display: flex; justify-content: flex-start; align-items: center;
  padding: 16px 24px; border-top: 1px solid #2a2a2a;
  background-color: #1E1E1E; flex-shrink: 0; gap: 16px;
}
.footer-action-btn {
  display: flex; align-items: center; gap: 8px;
  background-color: transparent; color: #D1CFC0; border: 1px solid #4a4a4a;
  padding: 8px 16px; border-radius: 20px; cursor: pointer;
  font-size: 14px; font-weight: 600;
  transition: all 0.2s ease;
}
.footer-action-btn:hover {
  background-color: rgba(247, 111, 83, 0.1);
  border-color: #F79253;
  color: #F79253;
}
.decrypt-container {
  margin-left: auto;
}
.decrypt-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  background-color: #F79253;
  color: #1E1E1E; /* Dark text for contrast */
  border: none;
  padding: 10px 20px;
  border-radius: 8px;
  cursor: pointer;
  font-size: 16px;
  font-weight: 600;
  transition: background-color 0.2s ease, transform 0.1s ease;
}
.decrypt-btn:hover {
  background-color: #e65a40;
}
.decrypt-btn:active {
  transform: scale(0.98);
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0,0,0, 0.6);
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
  max-width: 420px;
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

/* NEW: Style for the key input field */
.key-input {
  width: 100%;
  padding: 12px 16px;
  border: 1px solid #4a4a4a;
  background-color: #121212;
  color: #D1CFC0;
  border-radius: 8px;
  font-size: 16px;
  margin-bottom: 24px;
  box-sizing: border-box;
  outline: none;
  transition: border-color 0.2s, box-shadow 0.2s;
}

.key-input:focus {
  border-color: #F79253;
  box-shadow: 0 0 0 3px rgba(247, 111, 83, 0.2);
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
  background-color: #F79253;
  color: #1E1E1E;
}
.btn-confirm:hover {
  background-color: #e65a40;
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
  opacity: 0;
}
.fade-enter-from .confirm-dialog,
.fade-leave-to .confirm-dialog {
  transform: scale(0.95);
}
.btn-delete {
  background-color: #c0392b; /* A darker, more muted red */
  color: white;
}
.btn-delete:hover {
  background-color: #a93226;
}

.dialog-actions button:disabled {
  background-color: #555;
  color: #8A8A8A;
  cursor: not-allowed;
}
</style>
