<template>
    <div class="email-view-container">
    <div class="email-view-header">
      <button @click="$emit('close')" class="back-btn" title="Back to Inbox"> <i class="fa-solid fa-arrow-left"></i> </button> <div class="header-actions">
        <button class="options-btn" title="Archive"><i class="fa-solid fa-box-archive"></i></button>
        <button class="options-btn" title="Delete" @click="handleDelete(props.email)"><i class="fa-solid fa-trash"></i></button>
        <button class="options-btn" title="More options"><i class="fa-solid fa-ellipsis-v"></i></button>
        <div class="decrypt-container">
        <button v-if="props.email.isEncrypted" class="decrypt-btn" @click="openDecryptDialog">
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
import { DownloadAttachment, Decrypt, DeleteSingleEmail } from '../../wailsjs/go/main/App';
const showDecryptDialog = ref(false);
const decryptionKey = ref('');
const showDeleteDialog = ref(false);
const secret_key = ref('1234');

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
  background-color: #fff;
  font-family: 'Segoe UI', sans-serif;
}

/* Header */
.email-view-header {
  display: flex;
  align-items: center;
  padding: 8px 16px;
  border-bottom: 1px solid #eef1f5;
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
  color: #5f6368;
  font-size: 18px; 
  width: 40px; 
  height: 40px; 
  border-radius: 50%;
  display: flex; 
  align-items: center; 
  justify-content: center;
  transition: background-color 0.2s ease;
}
.back-btn:hover, .options-btn:hover { background-color: #f1f3f4; }

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
  flex-direction: row; /* Stacks the label and value vertically */
  align-items: flex; /* Aligns items to the left */
  border-bottom: 1px solid #eef1f5;
  padding: 12px 0; /* Adjusted padding for vertical layout */

}

/* The label is now styled as a small heading for the field below it */
.form-group label {
  font-size: 13px;
  color: #5f6368; /* Darker for better readability */
  font-weight: 600; /* Bolder */
  margin-bottom: 4px; /* Space between label and value */
  /* The fixed width and margin-right are removed */
}



/* This is the styled text that displays the value */
.value-display {
  flex-grow: 1;
  font-size: 15px;
  padding: 0; /* Padding is now on the parent */
  color: #2c3e50;
  text-align: left;
  padding-left: 10px;
  margin-bottom: auto;
}

.subject-value {
  font-size: 16px;
}

/* Main Email Content */
.email-body-content {
  flex-grow: 1;
  padding: 20px 0;
  font-size: 16px;
  line-height: 1.7;
  color: #34495e;
  white-space: pre-wrap;
  text-align: left;
}

/* Attachments */
.attachments-container {
  display: flex; flex-wrap: wrap; gap: 10px;
  padding: 16px 0; border-top: 1px solid #eef1f5;
}
.attachment-pill {
  display: flex; align-items: center; gap: 8px;
  background-color: #f4f6f8; border: 1px solid #eef1f5;
  color: #566573; font-size: 13px; font-weight: 500;
  border-radius: 16px; padding: 6px 12px;
}
.download-btn {
  background: none; border: none; cursor: pointer; color: #566573;
  transition: color 0.2s;
}
.download-btn:hover { color: #3498db; }

/* Footer */
.email-view-footer {
  display: flex; justify-content: flex-start; align-items: center;
  padding: 16px 24px; border-top: 1px solid #eef1f5;
  background-color: #f9fafb; flex-shrink: 0; gap: 16px;
}
.footer-action-btn {
  display: flex; align-items: center; gap: 8px;
  background-color: #ffffff; color: #5f6368; border: 1px solid #dadce0;
  padding: 8px 16px; border-radius: 20px; cursor: pointer;
  font-size: 14px; font-weight: 600;
  transition: all 0.2s ease;
}
.footer-action-btn:hover {
  background-color: #f8f9fa;
  border-color: #c6c6c6;
  box-shadow: 0 1px 2px rgba(0,0,0,.05);
}
.decrypt-container {
    margin-left: auto;
}
.decrypt-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  background-color: #8575db;  
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 8px;
  cursor: pointer;
  font-size: 15px;
  font-weight: 600;
  transition: background-color 0.2s ease, transform 0.1s ease;
}
.decrypt-btn:hover {
  background-color: #9775db;
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
  background-color: rgba(0,0,0, 0.4);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.confirm-dialog {
  background: #ffffff;
  padding: 32px;
  border-radius: 16px;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.2);
  width: 100%;
  max-width: 420px; /* Slightly wider for the input */
  text-align: center;
  transform: scale(1);
}

.confirm-dialog h3 {
  margin-top: 0;
  margin-bottom: 12px;
  font-size: 20px;
  color: #34495e;
}

.confirm-dialog p {
  margin-bottom: 24px;
  color: #7f8c8d;
  font-size: 15px;
  line-height: 1.6;
}

/* NEW: Style for the key input field */
.key-input {
  width: 100%;
  padding: 12px 16px;
  border: 1px solid #dcdcdc;
  border-radius: 8px;
  font-size: 16px;
  margin-bottom: 24px;
  box-sizing: border-box; /* Important for padding */
  outline: none;
  transition: border-color 0.2s, box-shadow 0.2s;
}

.key-input:focus {
  border-color: #3c5aed;
  box-shadow: 0 0 0 3px rgba(60, 90, 237, 0.2);
}

.dialog-actions {
  display: flex;
  justify-content: flex-end; /* Align buttons to the right */
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
  background-color: #ecf0f1;
  color: #7f8c8d;
}
.btn-cancel:hover {
  background-color: #e1e5e6;
}

.btn-confirm {
  background-color: #3c5aed;
  color: white;
}
.btn-confirm:hover {
  background-color: #344dc5;
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
  background-color: #d9534f; /* A standard destructive red */
  color: white;
}
.btn-delete:hover {
  background-color: #c9302c; /* A darker red on hover */
}
</style>
