<script lang="ts">
  import Header from "$lib/Header.svelte";
  import Sidebar from "$lib/Sidebar.svelte";
  import Login from "../lib/Login.svelte";
  import { currentUser, pb } from "$lib/pocketbase";
    import { onMount } from "svelte";

  onMount(()=>{
    if (!pb.authStore.model?.id){
      
    }
  })
</script>

{#if !$currentUser}
  <Login />
{:else if $currentUser}
  <Header />
  <div class="flex flex-col w-full lg:flex-row">
      <Sidebar />
    <div class="divider lg:divider-horizontal" />
    <div class="w-3/4 p-5 shadow-xl">
      <slot />
    </div>
  </div>
{/if}
