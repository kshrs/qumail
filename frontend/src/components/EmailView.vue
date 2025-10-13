<template>
    <div class="email-view-container">
    <div class="email-view-header">
      <button @click="$emit('close')" class="back-btn" title="Back to Inbox"> <i class="fa-solid fa-arrow-left"></i> </button> <div class="header-actions">
        <button class="options-btn" title="Archive"><i class="fa-solid fa-box-archive"></i></button>
        <button class="options-btn" title="Delete"><i class="fa-solid fa-trash"></i></button>
        <button class="options-btn" title="More options"><i class="fa-solid fa-ellipsis-v"></i></button>
        <div class="decrypt-container">
        <button v-if="props.email.isEncrypted" class="decrypt-btn" @click="decryptAction(props.email)">
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
    </div>

  </div>
</template>

<script setup>
import { ref } from 'vue';
import { DownloadAttachment, Decrypt } from '../../wailsjs/go/main/App';

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

defineEmits(['close']);

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
</style>
