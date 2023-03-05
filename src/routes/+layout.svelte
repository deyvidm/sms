<script lang="ts">
  import Header from "$lib/Header.svelte";
  import Sidebar from "$lib/Sidebar.svelte";
  import Login from "../lib/Login.svelte";
  import { currentUser } from "$lib/pocketbase";

  Array.prototype.remove = function (from, to) {
    var rest = this.slice((to || from) + 1 || this.length);
    this.length = from < 0 ? this.length + from : from;
    return this.push.apply(this, rest);
  };
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
