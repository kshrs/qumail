<template>
  <div class="login-container">
    <div class="login-card">
      <h1 class="logo">QuMail</h1>
      <h2 class="prompt">Sign in</h2>
      <p class="sub-prompt">to continue to your account</p>
      
      <button class="google-signin-btn" @click="handleSignIn">
        <div class="google-icon">
            <svg viewBox="0 0 18 18" role="presentation" aria-hidden="true" focusable="false">
                <path d="M17.64 9.2c0-.64-.05-1.28-.16-1.9H9v3.48h4.84a4.14 4.14 0 0 1-1.79 2.72v2.26h2.92c1.7-1.56 2.68-3.88 2.68-6.56z" fill="#4285F4"></path>
                <path d="M9 18c2.43 0 4.47-.8 5.96-2.18l-2.92-2.26c-.8.54-1.84.86-3.04.86-2.34 0-4.32-1.58-5.03-3.7H.96v2.34A8.99 8.99 0 0 0 9 18z" fill="#34A853"></path>
                <path d="M3.97 10.71A5.4 5.4 0 0 1 3.68 9c0-.59.1-1.16.29-1.71V4.96H.96A8.99 8.99 0 0 0 0 9c0 1.45.35 2.82.96 4.04l3.01-2.33z" fill="#FBBC05"></path>
                <path d="M9 3.54c1.32 0 2.5.45 3.44 1.34l2.58-2.58A8.96 8.96 0 0 0 9 0C5.06 0 1.63 2.31.96 5.59l3.01 2.32C4.68 5.12 6.66 3.54 9 3.54z" fill="#EA4335"></path>
            </svg>
        </div>
        <span>Sign in with Google</span>
      </button>

      <button class="outlook-signin-btn" @click="handleOutlookSignIn">
        <div class="outlook-icon">
            <svg width="17" height="17" viewBox="0 0 23 23" role="img" aria-label="Microsoft Logo">
  
              <rect class="ms-red"    x="0"  y="0"  width="10.5" height="10.5" />
              <rect class="ms-green"  x="12.5" y="0"  width="10.5" height="10.5" />
              <rect class="ms-blue"   x="0"  y="12.5" width="10.5" height="10.5" />
              <rect class="ms-yellow" x="12.5" y="12.5" width="10.5" height="10.5" />
            </svg>
        </div>
        <span>Sign in with Microsoft</span>
      </button>

    </div>
  </div>
</template>

<script setup>
import { SignIn } from '../../wailsjs/go/main/App';

async function handleSignIn() {
  try {
    const result = await SignIn();
    alert(result);
    // TODO: Proceed to main app view
  } catch (err) {
    alert(`Sign in failed: ${err}`);
  }
}

// NEW: Placeholder function for Outlook sign-in
function handleOutlookSignIn() {
  console.log("Sign in with Microsoft clicked");
  alert("Outlook sign-in is not yet implemented.");
  // TODO: Call the backend function for Microsoft OAuth
}
</script>

<style scoped>
/* Main container to center the login card */
.login-container {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background-color: #121212; /* Zen: Primary Background */
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

/* The dark card holding the content */
.login-card {
  background-color: #1E1E1E; /* Zen: UI/Surface Background */
  padding: 48px 40px;
  border-radius: 8px;
  border: 1px solid #2a2a2a; /* Subtle border for depth */
  box-shadow: none; /* Remove light theme shadow */
  width: 100%;
  max-width: 420px;
  text-align: center;
}

/* App name styling */
.logo, .prompt, .sub-prompt {
  text-align: center;
}

.logo {
  font-size: 28px;
  font-weight: 700;
  color: #E1E1E1; /* Zen: Primary Text */
  margin-bottom: 8px;
}

.prompt {
  font-size: 24px;
  font-weight: 600;
  color: #E1E1E1; /* Zen: Primary Text */
  margin-bottom: 8px;
}

.sub-prompt {
  font-size: 15px;
  color: #8A8A8A; /* Zen: Secondary Text */
  margin-bottom: 32px;
}

/* --- Base Button Styles --- */
.google-signin-btn,
.outlook-signin-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  padding: 12px;
  border-radius: 4px;
  font-size: 16px;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.2s ease, border-color 0.2s ease;
  margin-top: 16px; /* Space between buttons */
  background-color: transparent;
  border: 1px solid #4A4A4A; /* Darker border for buttons */
  color: #E1E1E1; /* Zen: Primary Text */
}

.google-signin-btn:hover,
.outlook-signin-btn:hover {
  background-color: rgba(79, 178, 165, 0.1); /* Zen: Accent with transparency */
  border-color: #4FB2A5; /* Zen: Primary Accent */
}


/* We keep the original brand colors for the icons for recognition */
.google-icon {
  width: 20px;
  height: 20px;
  margin-right: 12px;
}

.outlook-icon {
  width: 20px;
  height: 20px;
  margin-right: 12px;
}

/* Remove individual button styles that are now handled by the base class */
.google-signin-btn,
.outlook-signin-btn {
  /* These properties are now defined in the base style above */
}

.ms-red    { fill: #F25022; }
.ms-green  { fill: #7FBA00; }
.ms-blue   { fill: #00A4EF; }
.ms-yellow { fill: #FFB900; }
</style>

