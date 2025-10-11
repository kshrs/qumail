<template>
  <div class="email-view-container">
    <div class="email-view-header">
      <button @click="$emit('close')" class="back-btn" title="Back to Inbox">
        <i class="fa-solid fa-arrow-left"></i>
      </button>
      <div class="header-actions">
        <button class="options-btn" title="Archive"><i class="fa-solid fa-box-archive"></i></button>
        <button class="options-btn" title="Delete"><i class="fa-solid fa-trash"></i></button>
        <button class="options-btn" title="More options"><i class="fa-solid fa-ellipsis-v"></i></button>
      </div>
    </div>

    <div class="compose-body">
      <div class="form-group">
        <label>From</label>
        <div class="value-display">
          <strong>{{ props.email.sender }}</strong> &lt;{{ props.email.senderEmail }}&gt;
        </div>
      </div>
      
      <div class="form-group">
        <label>To</label>
        <div class="value-display">{{ props.email.to }}</div>
      </div>
      
      <div v-if="props.email.cc" class="form-group">
        <label>Cc</label>
        <div class="value-display">{{ props.email.cc }}</div>
      </div>

      <div class="form-group">
        <label>Subject</label>
        <div class="value-display subject-value">
          <strong>{{ props.email.subject }}</strong>
        </div>
      </div>
      
      <div class="email-body-content" v-html="props.email.body"></div>
      
      <div v-if="props.email.attachments && props.email.attachments.length > 0" class="attachments-container">
        <div v-for="(file, index) in props.email.attachments" :key="index" class="attachment-pill">
          <i class="fa-solid fa-paperclip"></i>
          <span>{{ file.name }} ({{ file.size }})</span>
          <button class="download-btn" title="Download">&darr;</button>
        </div>
      </div>
    </div>

    <div class="email-view-footer">
      <button class="footer-action-btn">
        <i class="fa-solid fa-reply"></i>
        <span>Reply</span>
      </button>
      <button class="footer-action-btn">
        <i class="fa-solid fa-reply-all"></i>
        <span>Reply All</span>
      </button>
      <button class="footer-action-btn">
        <i class="fa-solid fa-share"></i>
        <span>Forward</span>
      </button>
    </div>
  </div>
</template>

<script setup>

const props = defineProps({
  email: {
    type: Object,
    required: true,
  },
});

defineEmits(['close']);

const toggleStar = () => {
  props.email.isStarred = !props.email.isStarred;
};
</script>

<style src="../assets/styles/EmailView.css"></style>