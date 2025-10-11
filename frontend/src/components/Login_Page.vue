
<template>
    <div class="login-page">
        <div class="login-card" >
            <h2>Qumail</h2>

            <form @submit.prevent="onSubmit" class="login-form">
                <label class="field">
                    <span>Email:</span>
                    <input type="email" v-model="email" required placeholder="kabb@projectbase.com" />
                </label>

                <label class="field password-field">
                    <span>Password:</span>
                    <div class="password-row">
                        <input :type="showPassword ? 'text' : 'password'" v-model="password" required placeholder="Enter your password" />
                        <button
                            type="button"
                            class="icon-btn"
                            @click="showPassword = !showPassword"
                            :aria-pressed="showPassword"
                        >
                            <span class="material-symbols-outlined">{{ showPassword ? 'visibility_off' : 'visibility' }}</span>
                        </button>
                    </div>
                </label>

                <div class="actions-row">
                    <a href="#" @click.prevent="$emit('forgot-password')">Forgot password?</a>
                </div>

                <div class="btn">
                    <button type="submit" class="primary">Login</button>
                </div>
            </form>

            <div class="divider"><span>or</span></div>

            <button class="google-btn" @click="signInWithGoogle">
                <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/3/3c/Google_Favicon_2025.svg/768px-Google_Favicon_2025.svg.png?20250526093708"></img>
                <span>Sign in with Google</span>
            </button>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';
const emit = defineEmits(['login', 'forgot-password']);

const email = ref('');
const password = ref('');
const showPassword = ref(false);

function onSubmit() {
    emit('login', { email: email.value, password: password.value });
}

function signInWithGoogle() {
    const clientId = 'YOUR_GOOGLE_CLIENT_ID';
    const redirectUri = 'http://localhost:5173/auth/google/callback';
    const scope = encodeURIComponent('profile email');
    const url = `https://accounts.google.com/o/oauth2/v2/auth?response_type=code&client_id=${clientId}&redirect_uri=${encodeURIComponent(
        redirectUri
    )}&scope=${scope}`;
    window.open(url, 'google-oauth', 'width=500,height=600');

}

</script>

<style scoped>
.login-page {
    position: relative;
    height: calc(100vh - 56px);
    display: flex;
    align-items: center;
    justify-content: center;
    padding-top: 56px;
    overflow: hidden;
}

.login-card {
    position: relative;
    z-index: 3;
    width: 360px;
    background: rgba(255, 255, 255, 0.98);
    border-radius: 12px;
    box-shadow: 0 8px 30px rgba(0, 0, 0, 0.12);
    padding: 28px;
    display: flex;
    flex-direction: column;
    gap: 16px;
}
.login-card h2 {
    margin: 0;
    font-size: 20px;
    color:black;
}

.login-form {
    display: flex;
    flex-direction: column;
    gap: 12px;
}

.field span {
    display: flex;
    color: #6b7280;
    font-size: 13px;
    margin-bottom: 6px;
}

.field input {
    width: 100%;
    padding: 10px 12px;
    border: 1px solid #e5e7eb;
    border-radius: 8px;
    outline: none;
}

.password-row {
    display: flex;
    gap: 8px;
}

.icon-btn {
    background: transparent;
    border: none;
    color: #2563eb;
    cursor: pointer;
    padding: 6px 8px;
    border-radius: 6px;
}

.actions-row {
    display: flex;
    justify-content: center;
    align-items: center;
}

.actions-row a {
    color: #2563eb;
    text-decoration: none;
    font-size: 13px;
}

.primary {
    background: #2563eb;
    color: #fff;
    border: none;
    padding: 10px 16px;
    border-radius: 8px;
    cursor: pointer;
}


.divider span {
    background: rgba(255, 255, 255, 0.9);
    padding: 0 12px;
    color: #9ca3af;
}

.divider:before {
    content: '';
    height: 1px;
    background: #e5e7eb;
    flex: 1;
}

.google-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 10px;
    padding: 10px;
    border-radius: 8px;
    border: 1px solid #e5e7eb;
    background: #fff;
    cursor: pointer;
}

.google-btn img{
    height:20px;
}

.material-symbols-outlined {
    font-size: 18px !important;
    font-variation-settings: 'FILL' 0, 'wght' 400, 'GRAD' 0, 'opsz' 48;
}

.login-form .btn {
    width: 100%;
}
</style>