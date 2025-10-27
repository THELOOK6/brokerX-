<template>
  <div class="login-container">
    <div class="login-card">
      <h1 class="login-title">Sign In</h1>

      <el-form class="login-form" @submit.prevent="handleLogin">
        <!-- Email -->
        <el-form-item>
          <el-input
            v-model="email"
            placeholder="Email"
            prefix-icon="Message"
            type="email"
            autocomplete="username"
          />
        </el-form-item>

        <!-- Password -->
        <el-form-item>
          <el-input
            v-model="password"
            placeholder="Password"
            prefix-icon="Lock"
            show-password
            type="password"
            autocomplete="current-password"
          />
        </el-form-item>

        <!-- Login Button -->
        <el-button
          type="primary"
          class="login-btn"
          native-type="submit"
          round
        >
          Login
        </el-button>
      </el-form>

      <p class="signup-text">
        Don’t have an account?
        <RouterLink to="/signup" class="signup-link">Sign up</RouterLink>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Message, Lock } from '@element-plus/icons-vue'

const email = ref('')
const password = ref('')

const handleLogin = () => {
  if (!email.value || !password.value) {
    ElMessage.error('Please fill out all fields.')
    return
  }

  // Example login logic
  console.log('Logging in with', email.value, password.value)
  ElMessage.success('Login successful!')
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: #0f1117;
  color: #e5e7eb;
}

.login-card {
  background: #161a23;
  border: 1px solid #1b1d25;
  border-radius: 16px;
  padding: 2.5rem;
  width: 360px;
  text-align: center;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
}

.login-title {
  font-size: 2rem;
  font-weight: 700;
  color: #3b82f6;
  margin-bottom: 2rem;
}

/* Form */
.login-form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

/* Inputs */
.el-input__wrapper {
  background-color: #0f1117 !important;
  border: 1px solid #1b1d25 !important;
  box-shadow: none !important;
  color: #e5e7eb;
}
.el-input__inner {
  color: #e5e7eb;
}
.el-input__inner::placeholder {
  color: #6b7280;
}

/* Button */
.login-btn {
  width: 100%;
  background: linear-gradient(90deg, #2563eb, #3b82f6);
  border: none;
  color: white;
  font-weight: 600;
  height: 42px;
  margin-top: 1rem;
  box-shadow: 0 0 12px rgba(59, 130, 246, 0.3);
  transition: all 0.2s ease;
}
.login-btn:hover {
  box-shadow: 0 0 18px rgba(59, 130, 246, 0.5);
}

/* Signup */
.signup-text {
  margin-top: 1.5rem;
  font-size: 0.9rem;
  color: #9ca3af;
}
.signup-link {
  color: #3b82f6;
  text-decoration: none;
  margin-left: 4px;
}
.signup-link:hover {
  text-decoration: underline;
}
</style>
