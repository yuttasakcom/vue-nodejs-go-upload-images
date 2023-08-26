<template>
  <main>
    <div>
      <Header />
      <div class="mx-auto max-w-screen-2xl p-4 md:p-6 2xl:p-10">
        <div class="mx-auto max-w-270">
          <div class="col-span-5 xl:col-span-2">
            <div
              class="rounded-sm border border-stroke bg-white shadow-default dark:border-strokedark dark:bg-boxdark"
            >
              <div class="border-b border-stroke py-4 px-7 dark:border-strokedark">
                <h3 class="font-medium text-black dark:text-white">Your Photo</h3>
              </div>
              <div class="p-7">
                <form @submit.prevent="submit">
                  <div class="mb-5">
                    <label for="id_card_number" class="text-sm text-blue-700">
                      เลขบัตรประจำตัวประชาชน
                    </label>
                    <div>
                      <input
                        id="id_card_number"
                        v-model="id_card_number"
                        class="border rounded border-brand-gray-1 p-1 w-80"
                        type="text"
                      />
                    </div>
                  </div>

                  <div class="mb-5">
                    <label for="fileupload" class="text-sm text-blue-700">
                      Upload รูปภาพบัตรประชาชน
                    </label>

                    <div
                      id="FileUpload"
                      class="relative mb-5.5 block w-full cursor-pointer appearance-none rounded border-2 border-dashed border-primary bg-gray py-4 px-4 dark:bg-meta-4 sm:py-7.5"
                    >
                      <input
                        id="fileupload"
                        type="file"
                        accept="image/*"
                        class="absolute inset-0 z-50 m-0 h-full w-full cursor-pointer p-0 opacity-0 outline-none"
                        @change="upload($event)"
                      />
                      <input v-model="file" name="file" hidden />
                      <div class="flex flex-col items-center justify-center space-y-3">
                        <span
                          class="flex h-10 w-10 items-center justify-center rounded-full border border-stroke bg-white dark:border-strokedark dark:bg-boxdark"
                        >
                          <svg
                            width="16"
                            height="16"
                            viewBox="0 0 16 16"
                            fill="none"
                            xmlns="http://www.w3.org/2000/svg"
                          >
                            <path
                              fill-rule="evenodd"
                              clip-rule="evenodd"
                              d="M1.99967 9.33337C2.36786 9.33337 2.66634 9.63185 2.66634 10V12.6667C2.66634 12.8435 2.73658 13.0131 2.8616 13.1381C2.98663 13.2631 3.1562 13.3334 3.33301 13.3334H12.6663C12.8431 13.3334 13.0127 13.2631 13.1377 13.1381C13.2628 13.0131 13.333 12.8435 13.333 12.6667V10C13.333 9.63185 13.6315 9.33337 13.9997 9.33337C14.3679 9.33337 14.6663 9.63185 14.6663 10V12.6667C14.6663 13.1971 14.4556 13.7058 14.0806 14.0809C13.7055 14.456 13.1968 14.6667 12.6663 14.6667H3.33301C2.80257 14.6667 2.29387 14.456 1.91879 14.0809C1.54372 13.7058 1.33301 13.1971 1.33301 12.6667V10C1.33301 9.63185 1.63148 9.33337 1.99967 9.33337Z"
                              fill="#3C50E0"
                            />
                            <path
                              fill-rule="evenodd"
                              clip-rule="evenodd"
                              d="M7.5286 1.52864C7.78894 1.26829 8.21106 1.26829 8.4714 1.52864L11.8047 4.86197C12.0651 5.12232 12.0651 5.54443 11.8047 5.80478C11.5444 6.06513 11.1223 6.06513 10.8619 5.80478L8 2.94285L5.13807 5.80478C4.87772 6.06513 4.45561 6.06513 4.19526 5.80478C3.93491 5.54443 3.93491 5.12232 4.19526 4.86197L7.5286 1.52864Z"
                              fill="#3C50E0"
                            />
                            <path
                              fill-rule="evenodd"
                              clip-rule="evenodd"
                              d="M7.99967 1.33337C8.36786 1.33337 8.66634 1.63185 8.66634 2.00004V10C8.66634 10.3682 8.36786 10.6667 7.99967 10.6667C7.63148 10.6667 7.33301 10.3682 7.33301 10V2.00004C7.33301 1.63185 7.63148 1.33337 7.99967 1.33337Z"
                              fill="#3C50E0"
                            />
                          </svg>
                        </span>
                        <p class="font-medium text-sm">
                          <span class="text-primary">Click to upload</span>
                          or drag and drop
                        </p>
                        <p class="mt-1.5 font-medium text-sm">SVG, PNG, JPG or GIF</p>
                        <p class="font-medium text-sm">(max, 800 X 800px)</p>
                      </div>
                    </div>
                  </div>

                  <div class="mb-5">
                    <label for="bank_number" class="text-sm text-blue-700"> เลขบัญชีธนาคาร </label>
                    <div>
                      <input
                        id="bank_number"
                        v-model="bank_number"
                        class="border rounded border-brand-gray-1 p-1 w-80"
                        type="text"
                      />
                    </div>
                  </div>

                  <div class="mb-5">
                    <label for="fileupload" class="text-sm text-blue-700">
                      Upload รูปภาพบัญชีธนาคาร
                    </label>

                    <div
                      id="FileUpload"
                      class="relative mb-5.5 block w-full cursor-pointer appearance-none rounded border-2 border-dashed border-primary bg-gray py-4 px-4 dark:bg-meta-4 sm:py-7.5"
                    >
                      <input
                        id="fileupload2"
                        type="file"
                        accept="image/*"
                        count="2"
                        class="absolute inset-0 z-50 m-0 h-full w-full cursor-pointer p-0 opacity-0 outline-none"
                        @change="upload2($event)"
                      />
                      <input v-model="file2" name="file2" hidden />
                      <div class="flex flex-col items-center justify-center space-y-3">
                        <span
                          class="flex h-10 w-10 items-center justify-center rounded-full border border-stroke bg-white dark:border-strokedark dark:bg-boxdark"
                        >
                          <svg
                            width="16"
                            height="16"
                            viewBox="0 0 16 16"
                            fill="none"
                            xmlns="http://www.w3.org/2000/svg"
                          >
                            <path
                              fill-rule="evenodd"
                              clip-rule="evenodd"
                              d="M1.99967 9.33337C2.36786 9.33337 2.66634 9.63185 2.66634 10V12.6667C2.66634 12.8435 2.73658 13.0131 2.8616 13.1381C2.98663 13.2631 3.1562 13.3334 3.33301 13.3334H12.6663C12.8431 13.3334 13.0127 13.2631 13.1377 13.1381C13.2628 13.0131 13.333 12.8435 13.333 12.6667V10C13.333 9.63185 13.6315 9.33337 13.9997 9.33337C14.3679 9.33337 14.6663 9.63185 14.6663 10V12.6667C14.6663 13.1971 14.4556 13.7058 14.0806 14.0809C13.7055 14.456 13.1968 14.6667 12.6663 14.6667H3.33301C2.80257 14.6667 2.29387 14.456 1.91879 14.0809C1.54372 13.7058 1.33301 13.1971 1.33301 12.6667V10C1.33301 9.63185 1.63148 9.33337 1.99967 9.33337Z"
                              fill="#3C50E0"
                            />
                            <path
                              fill-rule="evenodd"
                              clip-rule="evenodd"
                              d="M7.5286 1.52864C7.78894 1.26829 8.21106 1.26829 8.4714 1.52864L11.8047 4.86197C12.0651 5.12232 12.0651 5.54443 11.8047 5.80478C11.5444 6.06513 11.1223 6.06513 10.8619 5.80478L8 2.94285L5.13807 5.80478C4.87772 6.06513 4.45561 6.06513 4.19526 5.80478C3.93491 5.54443 3.93491 5.12232 4.19526 4.86197L7.5286 1.52864Z"
                              fill="#3C50E0"
                            />
                            <path
                              fill-rule="evenodd"
                              clip-rule="evenodd"
                              d="M7.99967 1.33337C8.36786 1.33337 8.66634 1.63185 8.66634 2.00004V10C8.66634 10.3682 8.36786 10.6667 7.99967 10.6667C7.63148 10.6667 7.33301 10.3682 7.33301 10V2.00004C7.33301 1.63185 7.63148 1.33337 7.99967 1.33337Z"
                              fill="#3C50E0"
                            />
                          </svg>
                        </span>
                        <p class="font-medium text-sm">
                          <span class="text-primary">Click to upload</span>
                          or drag and drop
                        </p>
                        <p class="mt-1.5 font-medium text-sm">SVG, PNG, JPG or GIF</p>
                        <p class="font-medium text-sm">(max, 800 X 800px)</p>
                      </div>
                    </div>
                  </div>

                  <div class="flex justify-end gap-4.5">
                    <button
                      class="flex justify-center rounded mt-2 bg-brand-blue-1 py-2 px-6 font-medium text-white hover:bg-opacity-90"
                      type="submit"
                    >
                      Upload
                    </button>
                  </div>
                </form>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import Header from '@/components/Header.vue'

const id_card_number = ref<string>('')
const bank_number = ref<string>('')
let file = ref<File>()
let file2 = ref<File>()

const upload = (event: Event) => {
  const target = event.target as HTMLInputElement
  const files = target.files as FileList
  file.value = files[0]
}

const upload2 = (event: Event) => {
  const target = event.target as HTMLInputElement
  const files = target.files as FileList
  file2.value = files[0]
}

const submit = async () => {
  const formData = new FormData()
  if (id_card_number.value) {
    formData.append('id_card_number', id_card_number.value)
  }
  if (file.value) {
    formData.append('id_card_image', file.value)
  }
  if (bank_number.value) {
    formData.append('bank_number', bank_number.value)
  }
  if (file2.value) {
    formData.append('bank_image', file2.value)
  }

  const { data } = await axios.post('http://localhost:3000/upload', formData)
  console.log('file ====>', data)
}
</script>
