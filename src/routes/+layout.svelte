<script lang="ts">
  import Header from "$lib/Header.svelte";
  import Sidebar from "$lib/Sidebar.svelte";
  import Login from "../lib/Login.svelte";

  import { onMount } from "svelte"

  import { currentUser } from "$lib/gin";

  let savestore = false
  $: if (savestore && $currentUser) {
    window.sessionStorage.setItem("store", JSON.stringify($currentUser))
  }
  onMount(async () => {
    let ses = window.sessionStorage.getItem("store")
      if (ses) {
        console.log("sob-- ~ loading ses", ses)
        $currentUser = JSON.parse(ses)
      }
    savestore = true
  })

</script>

{#if $currentUser}
  <Header />
  <div class="flex flex-col w-full lg:flex-row">
    <Sidebar />
    <div class="divider lg:divider-horizontal" />
    <div class="w-3/4 p-5 shadow-xl">
      <slot />
    </div>
  </div>
{:else}
  <Login/>
{/if}
