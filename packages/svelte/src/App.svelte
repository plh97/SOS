<script lang="ts">
  import {
    Table,
    TableBody,
    TableBodyCell,
    TableBodyRow,
    TableHead,
    TableHeadCell,
  } from "flowbite-svelte";
  import "./app.css";
  import AddModal from "./components/AddModal.svelte";
  import type { Item } from "./interface";

  const tableData = $state<Item[]>([
    { position: "0", name: "Iron", symbol: "Fe", atomic_no: "26" },
    { position: "1", name: "Rhodium", symbol: "Rh", atomic_no: "45" },
    { position: "2", name: "Iodine", symbol: "I", atomic_no: "53" },
    { position: "3", name: "Radon", symbol: "Rn", atomic_no: "86" },
    { position: "4", name: "Technetium", symbol: "Tc", atomic_no: "43" },
  ]);

  const handleClick = (data: Item) => {
    tableData.push(data);
  };
</script>

<main class="p-2">
  <AddModal onSubmit={handleClick} />
  <Table>
    <TableHead>
      <TableHeadCell>Product name</TableHeadCell>
      <TableHeadCell>Color</TableHeadCell>
      <TableHeadCell>Category</TableHeadCell>
      <TableHeadCell>Price</TableHeadCell>
    </TableHead>
    <TableBody tableBodyClass="divide-y">
      {#each tableData as row}
        <TableBodyRow>
          <TableBodyCell>{row.position}</TableBodyCell>
          <TableBodyCell>{row.symbol}</TableBodyCell>
          <TableBodyCell>{row.name}</TableBodyCell>
          <TableBodyCell class="text-right">{row.atomic_no}</TableBodyCell>
        </TableBodyRow>
      {/each}
    </TableBody>
  </Table>
</main>

<style>
</style>
