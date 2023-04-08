<template>
  <v-container>
    <v-row class="pt-12" justify="center">
      <v-col cols="12" sm="12" md="8">
        <v-card variant="tonal">
          <v-card-title>
            <span class="headline">Register</span>
          </v-card-title>
          <v-card-text class="pb-0">
            <v-form ref="form" v-model="valid" lazy-validation>
              <v-text-field v-model="userForm.username" label="Username" required
                :rules="[rules.required]"></v-text-field>
              <v-text-field v-model="userForm.name" label="Name" required :rules="[rules.required]"></v-text-field>
              <v-text-field v-model="userForm.email" label="E-mail" required
                :rules="[rules.required, rules.email]"></v-text-field>
              <v-text-field v-model="userForm.password"
                :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'" :type="showPassword ? 'text' : 'password'"
                label="Password" hint="At least 10 characters" counter required
                @click:append="showPassword = !showPassword" :rules="[rules.min]"></v-text-field>
              <v-text-field v-model="userForm.passwordConfirm"
                :append-icon="showConfirmPassword ? 'mdi-eye' : 'mdi-eye-off'"
                :type="showConfirmPassword ? 'text' : 'password'" label="Confirm Password" hint="At least 10 characters"
                counter required @click:append="showConfirmPassword = !showConfirmPassword"
                :rules="[rules.min, rules.passwordMatch]"></v-text-field>
            </v-form>
          </v-card-text>
          <div class="pb-6 px-6">
            <v-btn size="x-large" variant="tonal" :disabled="!valid" color="success" @click="register" block>
              Submit
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
  min: (v: string) => v.length >= 10 || 'At least 10 characters',
  passwordMatch: () => (userForm.passwordConfirm === userForm.password) || 'Passwords must match',
  email: (v: string) => /^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(v) || 'E-mail must be valid',
};

const valid = ref(false);
const showPassword = ref(false);
const showConfirmPassword = ref(false);

const userForm = reactive({
  username: '',
  name: '',
  email: '',
  password: '',
  passwordConfirm: '',
});

async function register() {
  try {
    const registerSuccess = await appStore.registerUser(userForm)
    if (registerSuccess) {
      router.push('/login')
    }
  } catch (error) {
    if (error instanceof Error) {
      router.push('/n/error?msg=' + error.message)
    }
  }
}

</script>
