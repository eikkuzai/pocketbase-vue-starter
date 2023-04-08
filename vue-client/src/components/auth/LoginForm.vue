<template>
    <v-container>
      <v-row class="pt-12" justify="center">
        <v-col cols="12" sm="12" md="8">
          <v-card variant="tonal">
            <v-card-title>
              <span class="headline">Login</span>
            </v-card-title>
            <v-card-text class="pb-0">
              <v-form ref="form" v-model="valid" lazy-validation>
                <v-text-field v-model="loginForm.username" label="Username" required
                  :rules="[rules.required]"></v-text-field>
                <v-text-field v-model="loginForm.password"
                  :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'" :type="showPassword ? 'text' : 'password'"
                  label="Password" required
                  @click:append="showPassword = !showPassword"></v-text-field>
              </v-form>
            </v-card-text>
            <div class="pb-6 px-6">
              <v-btn size="x-large" variant="tonal" :disabled="!valid" color="success" @click="login" block>
                Login
              </v-btn>
            </div>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </template>
  
  <script setup lang="ts">
  import { reactive, ref } from 'vue';
  import { useAppStore } from '@/store/app';
  import router from '@/router';
  
  const appStore = useAppStore()
  
  const rules = {
    required: (v: string) => !!v || 'Required.',
    email: (v: string) => /^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(v) || 'E-mail must be valid',
  };
  
  const valid = ref(false);
  const showPassword = ref(false);
  
  const loginForm = reactive({
    username: '',
    password: ''
  });
  
  async function login() {
    try {
      const loginSuccess = await appStore.loginUser(loginForm)
      if (loginSuccess) {
        router.push('/')
      }
    } catch (error) {
      if (error instanceof Error) {
        router.push('/n/error?msg=' + error.message)
      }
    }
  }
  
  </script>
  