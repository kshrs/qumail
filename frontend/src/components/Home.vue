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
    </div>
  </div>
</template>

<script setup>
import { SignIn } from '../../wailsjs/go/main/App';

async function handleSignIn() {
  try {
    const result = await SignIn();
    alert(result); // "Successfully signed in!"
    // Now you can proceed to the main app view
    // Example: router.push('/inbox');
  } catch (err) {
    alert(`Sign in failed: ${err}`);
  }
}
</script>

<style scoped>
/* Main container to center the login card */
.login-container {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background-color: #f0f2f5; /* A light, neutral background */
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

/* The white card holding the content */
.login-card {
  background-color: #ffffff;
  padding: 48px 40px;
  border-radius: 8px;
  box-shadow: 0 4px 14px rgba(0, 0, 0, 0.09);
  width: 100%;
  max-width: 420px;
  text-align: left;
}

/* App name styling */
.logo {
  font-size: 28px;
  font-weight: 700;
  color: #1e1e1e;
  margin-bottom: 8px;
  text-align: center;
}

/* "Sign in" prompt */
.prompt {
  font-size: 24px;
  font-weight: 600;
  color: #333;
  margin-bottom: 8px;
  text-align: center;
}

.sub-prompt {
  font-size: 15px;
  color: #555;
  margin-bottom: 32px; /* Adds space before the button */
  text-align: center;
}

/* Google Sign-in Button */
.google-signin-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  padding: 12px;
  border: 1px solid #dadce0;
  border-radius: 4px;
  background-color: #ffffff;
  font-size: 16px;
  font-weight: 500;
  color: #3c4043;
  cursor: pointer;
  transition: background-color 0.2s ease, box-shadow 0.2s ease;
}

.google-signin-btn:hover {
  background-color: #f8f9fa; /* Slight gray on hover */
  border-color: #c6c9cc;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}

/* START: New Styles for the Icon */
.google-icon {
  width: 20px;
  height: 20px;
  margin-right: 12px; /* Space between icon and text */
  display: flex; /* Helps center the SVG if needed */
  align-items: center;
  justify-content: center;
}

/* This targets the SVG element itself inside the div */
.google-icon svg {
  width: 100%;
  height: 100%;
}
</style>
