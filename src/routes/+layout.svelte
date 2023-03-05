<script lang="ts">
  import Header from "$lib/Header.svelte";
  import Sidebar from "$lib/Sidebar.svelte";
  import Login from "../lib/Login.svelte";
  import { currentUser } from "$lib/pocketbase";
</script>

{#if !$currentUser}
  <Login />
{:else if $currentUser}
  <Header />
  <div class="flex flex-col w-full lg:flex-row">
    <Sidebar />
    <div class="divider lg:divider-horizontal" />
    <div class=" w-5/12 p-5 shadow-xl">
      <slot />
    </div>
  </div>
{/if}
