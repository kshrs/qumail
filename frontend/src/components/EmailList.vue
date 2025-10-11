<template>
  <div class="inbox-container">
    <div class="inbox-toolbar">
      </div>

    <div class="email-list">
      <div
        v-for="email in emails"
        :key="email.id"
        class="email-row"
        :class="{
          'is-read': email.isRead,
          'is-selected': email.isSelected,
        }"
        @click="handleRowClick(email)"
      >
        <div class="email-actions">
          <input
            type="checkbox"
            :checked="email.isSelected"
            @click.stop="toggleSelect(email)"
            class="email-checkbox"
          />
          <button @click.stop="toggleStar(email)" class="star-btn">
            <i :class="email.isStarred ? 'fa-solid fa-star' : 'fa-regular fa-star'"></i>
          </button>
        </div>
        
        <div class="email-sender">
          <span>{{ email.sender }}</span>
        </div>
        
        <div class="email-content">
          <span class="email-subject">{{ email.subject }}</span>
          <span class="email-snippet">&nbsp;-&nbsp;{{ email.snippet }}</span>
        </div>

        <div class="email-date">
          {{ email.timestamp.split(',')[0] }} </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const emit = defineEmits(['view-email']);


// This ref now contains the email objects, which will make them appear on the screen.
const emails = ref([
  {
    id: 1,
    sender: 'GitHub',
    senderEmail: 'noreply@github.com',
    subject: 'New vulnerability found in your repository',
    snippet: 'A high-severity vulnerability has been detected in the `express` dependency...',
    body: `<p>Hi DevTeam,</p><p>We've detected a high-severity vulnerability in the <strong>express</strong> dependency in your 'Project-Phoenix' repository.</p><p>We recommend updating to the latest patched version as soon as possible. Please review the security advisory for more details.</p><p>Thanks,<br>The GitHub Security Team</p>`,
    timestamp: 'Oct 10, 2025, 06:45 PM',
    isRead: false,
    isStarred: true,
    isSelected: false,
    to: 'you@example.com',
    cc: 'security@example.com',
    attachments: [{ name: 'vulnerability-report.pdf', size: '256 KB' }],
  },
  {
    id: 2,
    sender: 'Alice Johnson',
    senderEmail: 'alice.j@workplace.com',
    subject: 'Project Kick-off Meeting',
    snippet: 'Hi team, please find the agenda for our kick-off meeting tomorrow morning.',
    body: `<p>Hi Team,</p><p>Looking forward to our meeting tomorrow at 9:00 AM. I've attached the agenda for the Project Nova kick-off.</p><p>Please come prepared with any initial questions.</p><p>Best,<br>Alice</p>`,
    timestamp: 'Oct 10, 2025, 04:10 PM',
    isRead: false,
    isStarred: false,
    isSelected: false,
    to: 'you@example.com',
    cc: 'manager@example.com',
    attachments: [{ name: 'Project-Nova-Agenda.docx', size: '78 KB' }],
  },
  {
    id: 3,
    sender: 'Linear',
    senderEmail: 'notifications@linear.app',
    subject: 'PRJ-124: New comment on "Deploy frontend to production"',
    snippet: 'Bob commented: "Looks good to me! Ready for deployment."',
    body: `<p><strong>Bob Williams</strong> left a comment on issue <a href="#">PRJ-124</a>:</p><blockquote>"Looks good to me! Ready for deployment. Let's sync up before pushing the button."</blockquote>`,
    timestamp: 'Oct 10, 2025, 02:30 PM',
    isRead: true,
    isStarred: false,
    isSelected: false,
    to: 'you@example.com',
    cc: null,
    attachments: [],
  },
  {
    id: 4,
    sender: 'Figma',
    senderEmail: 'team@figma.com',
    subject: 'Weekly Design Updates & Inspiration',
    snippet: 'Check out the latest features, plugins, and stunning designs from the community.',
    body: `<p>Hi there,</p><p>Here's your weekly dose of inspiration from the Figma community. We've also rolled out some new features for auto-layout that you might find interesting!</p><p>Happy designing!</p>`,
    timestamp: 'Oct 09, 2025, 11:00 AM',
    isRead: true,
    isStarred: false,
    isSelected: false,
    to: 'you@example.com',
    cc: null,
    attachments: [],
  },
  {
    id: 5,
    sender: 'Charlie Davis',
    senderEmail: 'charlie.d@workplace.com',
    subject: 'Quick question about the API docs',
    snippet: "Hey, are the docs for the new payment gateway endpoint updated yet?",
    body: `<p>Hey,</p><p>Just checking in to see if the documentation for the new payment gateway endpoint is live. Can't seem to find it on Confluence.</p><p>Let me know!</p><p>Thanks,<br>Charlie</p>`,
    timestamp: 'Oct 08, 2025, 05:22 PM',
    isRead: true,
    isStarred: true,
    isSelected: false,
    to: 'you@example.com',
    cc: null,
    attachments: [],
  }
]);


const handleRowClick = (email) => {
  email.isRead = true; // Mark the email as read
  emit('view-email', email); // Send the email data to the parent component
};

// These functions handle the checkbox and star actions without navigating away
const toggleSelect = (email) => {
  email.isSelected = !email.isSelected;
};

const toggleStar = (email) => {
  email.isStarred = !email.isStarred;
};
</script>

<style src="../assets/styles/EmailList.css"></style>