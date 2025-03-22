<script lang="ts">
  import {
    Button,
    Modal,
    Label,
    Input,
    Select,
    Datepicker,
  } from "flowbite-svelte";

  let formModal = $state(false);
  let { onSubmit } = $props();
  let selectedDate = $state(new Date()); // 存储 Datepicker 的值
  const handleSubmit = (e: any) => {
    const data = new FormData(e.target);
    const json: Record<string, FormDataEntryValue | Date> = {
      enrollmentDate: selectedDate,
    };
    for (const [name, value] of data) {
      json[name] = value;
    }
    onSubmit(json);
    formModal = false;
  };

  let genderEnum = [
    { value: "male", name: "Male" },
    { value: "female", name: "Female" },
    { value: "other", name: "Other" },
  ];

  let statusEnum = [
    { value: "active", name: "Active" },
    { value: "inactive", name: "Inactive" },
    { value: "graduated", name: "Graduated" },
  ];
</script>

<Button class="mb-4" on:click={() => (formModal = true)}>Form modal</Button>

<Modal bind:open={formModal} size="xs" autoclose={false} class="w-full">
  <form
    class="flex flex-col space-y-6"
    action="#"
    on:submit|preventDefault={handleSubmit}
  >
    <h3 class="mb-4 text-xl font-medium text-gray-900 dark:text-white">
      Add Student
    </h3>

    <Label class="space-y-2">
      <span>Name</span>
      <Input name="name" placeholder="input name" />
    </Label>
    <Label class="space-y-2">
      <span>Email</span>
      <Input class="mt-2" type="email" name="email" placeholder="Input Email" />
    </Label>
    <Label class="space-y-2">
      <span>Gender</span>
      <Select class="mt-2" items={genderEnum} name="gender" />
    </Label>
    <Label class="space-y-2">
      <span>Grade</span>
      <Input class="mt-2" name="grade" placeholder="Input Grade" />
    </Label>
    <Label class="space-y-2">
      <span class="mb-2">EnrollmentDate</span>
      <Datepicker value={selectedDate} placeholder="Input EnrollmentDate" />
    </Label>
    <Label class="space-y-2">
      <span>Age</span>
      <Input class="mt-2" name="age" placeholder="Input Age" />
    </Label>
    <Label class="space-y-2">
      <span>Status</span>
      <Select
        class="mt-2"
        placeholder="Select Status"
        items={statusEnum}
        name="status"
      />
    </Label>
    <Button type="submit" class="w-full">Add Student</Button>
  </form>
</Modal>
