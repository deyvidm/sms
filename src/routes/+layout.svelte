<script lang="ts">
  import Header from "$lib/Header.svelte";
  import Sidebar from "$lib/Sidebar.svelte";
  import Login from "../lib/Login.svelte";

  import { onMount } from "svelte"

  import { currentUser, userEvents } from "$lib/gin";

  let savestore = false
  
  onMount(async () => {
    console.log("mounted root layout")
    // let ses = window.sessionStorage.setItem("store","")
    //   if (ses) {
    //     console.log("sob-- ~ loading ses", ses)
    //     $currentUser = JSON.parse(ses)
    //   }
    // savestore = true
  })

</script>

{#if $currentUser.username.length > 0}
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
