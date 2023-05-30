<script lang="ts">
  import ContactList from "$lib/ContactList.svelte";
  import { apiClient } from "$lib/gin";
  import type { Contact } from "$lib/gripes";
  import { onMount } from "svelte";

  /** @type {import('./$types').PageData} */
  export let data;

  let title: string;
  let invite_body: string;
  let available = data.contacts;
  let invited = new Array();

  let submitButton: HTMLButtonElement;

  function handleInvite(event) {
    console.log(available.length);
    available = available.filter((c) => c.id != event.detail.recipient.id);
    invited = [...invited, event.detail.recipient];
  }

  function handleUninvite(event) {
    console.log(event);
    invited = invited.filter((c) => c.id != event.detail.recipient.id);
    available = [...available, event.detail.recipient];
  }

  function extractIDs(iterable: Array<Contact>) {
    return iterable.map((c) => c.id);
  }

  async function create() {
    submitButton.classList.add("loading");
    const body = await apiClient
      .post("/events/new", {
        title: title,
        invite_body: invite_body,
        contacts: extractIDs(invited),
      })
      .then((body) => {
        submitButton.classList.remove("loading");
        submitButton.classList.add("btn-disabled");
        submitButton.classList.add("btn-primary");
        submitButton.textContent = "Success";
      });
  }

  onMount(async () => {});
</script>

<h2 class="mb-10 text-4xl font-extrabold dark:text-white">Create New Event</h2>

<div class="form-control w-full">
  <form class="space-y-4 md:space-y-6" on:submit|preventDefault={create}>
    <div>
      <label class="label" for="event-title">
        <span class="text-xl font-bold text-xllabel-text">Title</span>
      </label>
      <input
        bind:value={title}
        type="text"
        name="title"
        placeholder="Sunday Funday St Paddy's Vball"
        class="input input-bordered w-full max-w"
      />
    </div>
    <div class="mt-5">
      <label class="label" for="event-content">
        <span class="text-xl font-bold text-xllabel-text">SMS Content</span>
      </label>
      <textarea
        bind:value={invite_body}
        name="content"
        class="w-full h-96 pb-10 textarea textarea-bordered"
      />
    </div>

    <div class="flex space-x-4">
      <span>
        <span class="text-xl font-bold text-xllabel-text">Available</span>
        <ContactList contacts={available} on:message={handleInvite} />
      </span>
      <span>
        <span class="text-xl font-bold text-xllabel-text">Invited</span>
        <ContactList contacts={invited} on:message={handleUninvite} />
      </span>
    </div>
    <input type="hidden" name="recipients" value={invited} />
    {#if invited.length > 0}
      <button bind:this={submitButton} class="mt-12 w-1/4 btn btn-active"
        >Create</button
      >
    {/if}
  </form>
</div>
