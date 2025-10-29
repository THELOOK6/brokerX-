<template>
  <div class="signup-container">
    <div class="signup-card">
      <h1 class="signup-title">Sign Up</h1>

      <el-form class="signup-form" @submit.prevent="handleSignup">
        <!-- Full Name -->
        <el-form-item>
          <el-input
            v-model="name"
            placeholder="Full Name"
            prefix-icon="Message"
            type="name"
            autocomplete="full-name"
          />
        </el-form-item>
        
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
          class="signup-btn"
          native-type="submit"
          round
        >
          Sign Up
        </el-button>
      </el-form>

      <p class="login-text">
        Already have an account?
        <RouterLink to="/login" class="login-link">Login</RouterLink>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Message, Lock } from '@element-plus/icons-vue'

const name = ref('')
const email = ref('')
const password = ref('')

const API_BASE_URL = 'http://localhost:8080/api/v1/auth/signup' 

const handleSignup = async () => {
  if (!name.value || !email.value || !password.value) {
    ElMessage.error('Please fill out all fields.')
    return
  }

  try {
    const response = await fetch(API_BASE_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        email: email.value,
        password: password.value,
        full_name: name.value
      })
    })

    if (!response.ok) {
      // Try to read backend error message if any
      const errorData = await response.json().catch(() => ({}))
      const msg = errorData?.error || `Signup failed (${response.status})`
      throw new Error(msg)
    }

    const data = await response.json()
    console.log('Signup successful:', data)

    ElMessage.success(`Account created successfully for ${data.email}`)
    // Optional: redirect or clear fields
    name.value = ''
    email.value = ''
    password.value = ''
  } catch (error: any) {
    console.error('Signup error:', error)
    ElMessage.error(error.message || 'An unexpected error occurred.')
  }
}
</script>


<style scoped>
.signup-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: #0f1117;
  color: #e5e7eb;
}

.signup-card {
  background: #161a23;
  border: 1px solid #1b1d25;
  border-radius: 16px;
  padding: 2.5rem;
  width: 360px;
  text-align: center;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
}

.signup-title {
  font-size: 2rem;
  font-weight: 700;
  color: #3b82f6;
  margin-bottom: 2rem;
}

/* Form */
.signup-form {
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
.signup-btn {
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
.signup-btn:hover {
  box-shadow: 0 0 18px rgba(59, 130, 246, 0.5);
}

/* Login */
.login-text {
  margin-top: 1.5rem;
  font-size: 0.9rem;
  color: #9ca3af;
}
.login-link {
  color: #3b82f6;
  text-decoration: none;
  margin-left: 4px;
}
.login-link:hover {
  text-decoration: underline;
}
</style>
