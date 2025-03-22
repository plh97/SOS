<script lang="ts">
  import moment from "moment";
  import {
    Button,
    Table,
    TableBody,
    TableBodyCell,
    TableBodyRow,
    TableHead,
    TableHeadCell,
    Spinner, // 新增 Spinner 组件
  } from "flowbite-svelte";
  import "./app.css";
  import AddModal from "./components/AddModal.svelte";
  import type { Student } from "./interface";
  import { createStudent, deleteStudent, getStudentList } from "./Api";

  // 定义状态，包括 tableData 和 loading
  const state = $state<{
    tableData: Student[];
    loading: boolean; // 新增 loading 状态
  }>({
    tableData: [],
    loading: true, // 初始为加载中
  });

  // 获取学生列表，加载完成后更新状态
  getStudentList()
    .then((data) => {
      state.tableData = data;
      state.loading = false; // 数据加载完成，关闭加载状态
    })
    .catch((error) => {
      console.error("Failed to load students:", error);
      state.loading = false; // 出错时也关闭加载状态
    });

  // 添加学生
  const handleClick = async (data: Student) => {
    const res = await createStudent(data);
    state.tableData.push(res);
  };

  // 删除学生
  const handleDelete = async (id: string) => {
    await deleteStudent(id);
    state.tableData = state.tableData.filter((item) => item.id !== id);
  };
</script>

<main class="p-2">
  <AddModal onSubmit={handleClick} />

  {#if state.loading}
    <!-- 加载状态显示 -->
    <div class="flex flex-col justify-center items-center w-full h-12 mt-12">
      <Spinner size="xl" />
      <span class="text-gray-700">Loading...</span>
    </div>
  {:else}
    <!-- 数据加载完成后显示表格 -->
    <Table width="100%">
      <TableHead>
        <TableHeadCell>ID</TableHeadCell>
        <TableHeadCell>Name</TableHeadCell>
        <TableHeadCell>Email</TableHeadCell>
        <TableHeadCell>Gender</TableHeadCell>
        <TableHeadCell>Grade</TableHeadCell>
        <TableHeadCell>EnrollmentDate</TableHeadCell>
        <TableHeadCell>Age</TableHeadCell>
        <TableHeadCell>Status</TableHeadCell>
        <TableHeadCell class="text-right">Actions</TableHeadCell>
        <!-- 操作列 -->
      </TableHead>
      <TableBody tableBodyClass="divide-y">
        {#each state.tableData as row}
          <TableBodyRow>
            <TableBodyCell class="w-4">{row.id}</TableBodyCell>
            <TableBodyCell>{row.name}</TableBodyCell>
            <TableBodyCell>{row.email}</TableBodyCell>
            <TableBodyCell>{row.gender}</TableBodyCell>
            <TableBodyCell>{row.grade}</TableBodyCell>
            <TableBodyCell
              >{moment(row.enrollmentDate).format(
                "YYYY-MM-DD hh:mm:ss"
              )}</TableBodyCell
            >
            <TableBodyCell>{row.age}</TableBodyCell>
            <TableBodyCell>{row.status}</TableBodyCell>
            <TableBodyCell class="text-right">
              <Button on:click={() => handleDelete(row.id)}>Delete</Button>
            </TableBodyCell>
          </TableBodyRow>
        {/each}
      </TableBody>
    </Table>
  {/if}
</main>

<style>
  /* 可选：自定义加载状态样式 */
  .h-64 {
    height: 16rem; /* 确保加载时占用一定空间 */
  }
</style>
