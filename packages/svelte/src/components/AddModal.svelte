<script lang="ts">
  import { Button, Modal, Label, Input } from "flowbite-svelte";

  let formModal = $state(false);
  let { onSubmit } = $props();
  const handleSubmit = (e: any) => {
    const data = new FormData(e.target);
    const json: Record<string, FormDataEntryValue> = {};
    for (const [name, value] of data) {
      json[name] = value;
    }
    onSubmit(json);
    formModal = false;
  };
</script>

<Button class="mb-4" on:click={() => (formModal = true)}>Form modal</Button>

<Modal bind:open={formModal} size="xs" autoclose={false} class="w-full">
  <form
    class="flex flex-col space-y-6"
    action="#"
    on:submit|preventDefault={handleSubmit}
  >
    <h3 class="mb-4 text-xl font-medium text-gray-900 dark:text-white">
      Add data
    </h3>
    <Label class="space-y-2">
      <span>Name</span>
      <Input name="position" placeholder="input name" />
    </Label>
    <Label class="space-y-2">
      <span>Color</span>
      <Input name="symbol" placeholder="Input Color" />
    </Label>
    <Label class="space-y-2">
      <span>Category</span>
      <Input name="name" placeholder="Input Category" />
    </Label>
    <Label class="space-y-2">
      <span>Price</span>
      <Input name="atomic_no" placeholder="Input Price" />
    </Label>
    <Button type="submit" class="w-full">Add Data</Button>
  </form>
</Modal>
