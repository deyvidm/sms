<script lang="ts">
  import ContactList from "$lib/ContactList.svelte";
    import { apiClient } from "$lib/gin";

  let title: string;
  let invite_body: string;
  let recipients = new Array();
  let submitButton: HTMLButtonElement;
  
  function handleMessage(event) {
    recipients = event.detail.recipients;
  }

function extractIDs(iterable) {
  return Array.from(iterable, node => node.id);
}

  async function create() {
    console.log("create new event ---")
    submitButton.classList.add("loading")
    const d = new Date(Date.now());
    const body = await apiClient.post("/events/new",{
      title: title,
      invite_body: invite_body,
      contacts: extractIDs(recipients)
    }).then((body)=>{
        console.log(body)
        submitButton.classList.remove("loading")
        submitButton.classList.add("btn-disabled")
        submitButton.classList.add("btn-primary")
        submitButton.textContent = "Success"
    });

    // createEvent(<EventRecord>{
    //     organizer: pb.authStore.model?.id,
    //     title: formTitle,
    //     description: content,
    //     capacity: 666,
    //     start_date: d.toISOString(),
    //     end_date: d.toISOString(),
    //     send_invite_date: d.toISOString(),
    //     status: EventStatusOptions.active,
    // }).then((eventRecord) => {
    //     recipients.forEach((r) => {
    //         pb.collection('attendee')
    //             .create(<AttendeeRecord>{
    //                 event: eventRecord.id,
    //                 contact: r.id,
    //                 status: AttendeeStatusOptions['sending-invite'],
    //                 paid: false,
    //             })
    //             .catch((error) => {
    //                 console.log(error);
    //             });
    //     });
    // }).then(()=>{
    //     submitButton.classList.remove("loading")
    //     submitButton.classList.add("btn-disabled")
    //     submitButton.textContent = "Success"
    // });
  }
  /** @type {import('./$types').PageData} */
  export let data;
  let contacts = new Array();
  contacts = data.user.contacts;
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

    <ContactList {contacts} on:message={handleMessage} />
    <input type="hidden" name="recipients"  value={recipients}/>
    {#if recipients.length > 0}
      <button bind:this={submitButton}  class="mt-12 w-1/4 btn btn-active">Create</button>
    {/if}
  </form>
</div>
