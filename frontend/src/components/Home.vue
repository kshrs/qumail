<template>
  <div class="login-container">
    <h1 class="logo">QuMail</h1>
    <p class="logo-explanation">Welcome to <br>a <b class="word-secure"> secure </b> emailclient</p>
    <div class="login-card">
      <!-- Left Panel for Social Logins -->
      <div class="left-panel">
        <h2 class="prompt">Sign in</h2>
        <p class="sub-prompt">to continue to your account</p>

        <div class="input-group">
            <label for="email">Email Address</label>
            <input type="email" id="email" placeholder="your.name@.providercom" />
        </div>
        <button class="sso-btn">Sign in</button>
        <b style="font-family: Junicode, serif; font-size: 20px; margin-top: 10px;"> or </b>

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
              <rect class="ms-red" x="0" y="0" width="10.5" height="10.5" />
              <rect class="ms-green" x="12.5" y="0" width="10.5" height="10.5" />
              <rect class="ms-blue" x="0" y="12.5" width="10.5" height="10.5" />
              <rect class="ms-yellow" x="12.5" y="12.5" width="10.5" height="10.5" />
            </svg>
          </div>
          <span>Sign in with Microsoft</span>
        </button>
        <p class="sso-info">Use your service provider authentication for a seamless experience.</p>
      </div>

      <!-- Vertical Separator -->
      <div class="separator"></div>

      <!-- Right Panel for SSO Login -->
      <div class="right-panel">
        <h3 class="sso-prompt">Organizational Sign-in</h3>
        <div class="input-group">
            <label for="email">Email Address</label>
            <input type="email" id="email" placeholder="your.name@company.com" />
        </div>
        <button class="sso-btn">Sign in with SSO</button>
        <p class="sso-info">Use your organization's credentials for a seamless experience.</p>
      </div>
    </div>
  </div>
</template>

<script setup>
// Script remains the same as no new functionality was requested
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

function handleOutlookSignIn() {
  console.log("Sign in with Microsoft clicked");
  alert("Outlook sign-in is not yet implemented.");
  // TODO: Call the backend function for Microsoft OAuth
}
</script>

<style scoped>
/* Main container to center the logo and card */
.login-container {
  display: flex;
  flex-direction: column; /* Stack logo and card vertically */
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background-color: #121212;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

/* App name styling (outside the card) */
.logo {
  font-size: 6rem; /* Bigger font size */
  font-weight: 700;
  font-family: Junicode,serif;
  color: #E1E1E1;
  margin-bottom: 32px; /* Space between logo and card */
  text-align: center;
}

.logo-explanation {
  font-size: 4rem; /* Bigger font size */
  font-weight: 700;
  font-family: Junicode-italic,serif;
  color: #E1E1E1;
  margin-bottom: 32px; /* Space between logo and card */
  text-align: center;
}
.logo-explanation b {
  color: #F79253;
  text-align: center;
}
 
/* The dark card holding the content */
.login-card {
  background-color: #1E1E1E;
  padding: 40px;
  border-radius: 8px;
  border: 1px solid #2a2a2a;
  box-shadow: none;
  width: 100%;
  max-width: 800px; /* Increased max-width for two columns */
  display: flex; /* Enable flexbox for two-column layout */
  align-items: center;
}

/* --- Left & Right Panels --- */
.left-panel, .right-panel {
    flex: 1; /* Each panel takes up half the space */
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 0 20px;
    margin-bottom: auto;
}

/* Vertical separator line */
.separator {
    width: 1px;
    height: 100%;
    align-self: stretch; /* Make it full height of the card */
    background-color: #363636;
    margin: 0 20px;
}

/* --- Left Panel Content --- */
.prompt {
  font-size: 24px;
  font-weight: 600;
  color: #E1E1E1;
  margin-bottom: 8px;
}

.sub-prompt {
  font-size: 15px;
  color: #8A8A8A;
  margin-bottom: 32px;
}

/* --- Base Button Styles for Social Logins --- */
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
  margin-top: 16px;
  background-color: transparent;
  border: 1px solid #4A4A4A;
  color: #E1E1E1;
}

.google-signin-btn:hover,
.outlook-signin-btn:hover {
  background-color: rgba(245, 146, 83, 0.1);
  border-color: #F59253;
}

.google-icon, .outlook-icon {
  width: 20px;
  height: 20px;
  margin-right: 12px;
}

/* Microsoft logo colors */
.ms-red    { fill: #F25022; }
.ms-green  { fill: #7FBA00; }
.ms-blue   { fill: #00A4EF; }
.ms-yellow { fill: #FFB900; }


/* --- Right Panel Content --- */
.right-panel {
    text-align: center;
}
.sso-prompt {
    font-size: 24px;
    font-weight: 600;
    color: #E1E1E1;
    margin-bottom: 24px;
}
.input-group {
    width: 100%;
    text-align: left;
    margin-bottom: 20px;
}
.input-group label {
    display: block;
    color: #8A8A8A;
    font-size: 14px;
    margin-bottom: 8px;
}
.input-group input {
    width: 100%;
    padding: 12px;
    background-color: #121212;
    border: 1px solid #4A4A4A;
    color: #E1E1E1;
    border-radius: 4px;
    font-size: 16px;
    outline: none;
    transition: border-color 0.2s ease, box-shadow 0.2s ease;
}
.input-group input:focus {
    border-color: #f79253;
    box-shadow: 0 0 0 3px rgba(247, 146, 83, 0.2);
}
.sso-btn {
    width: 100%;
    padding: 12px;
    background-color: #f79253; /* New orange color */
    color: #1E1E1E; /* Dark text for contrast */
    border: none;
    border-radius: 4px;
    font-size: 16px;
    font-weight: 600;
    cursor: pointer;
    transition: background-color 0.2s ease;
}
.sso-btn:hover {
    background-color: #F76F53;
}
.sso-info {
    font-size: 12px;
    color: #8A8A8A;
    margin-top: 16px;
    line-height: 1.5;
}
</style>
