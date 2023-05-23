<script lang="ts">
  import Header from "./Header.svelte";
  import Sidebar from "$lib/Sidebar.svelte";
  import { navigating } from "$app/stores";

  import Login from "../lib/Login.svelte";

  import { onMount } from "svelte";

  import { currentUser, userEvents } from "$lib/gin";

  /** @type {import('./$types').PageData} */
  export let data;

  onMount(async () => {
    console.log("mounted root layout");
  });
</script>

<Header user={data.user}/>
{#if $navigating}
  <h1>We out here navigating</h1>
{/if}

{#if data.user}
<div class="flex flex-col w-full lg:flex-row">
  <Sidebar />
  <div class="divider lg:divider-horizontal" />
  <!-- <div class="w-3/4 p-5 shadow-xl"> -->
    <slot />
  <!-- </div> -->
</div>
{:else}
<slot />
{/if}
