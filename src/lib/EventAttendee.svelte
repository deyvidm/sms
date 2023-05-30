<script lang="ts">
    import { afterUpdate, createEventDispatcher, onMount } from 'svelte';
    import { attendees, get50Attendees, pb, updateAttendee } from './pocketbase';
    import type { AttendeeResponse, EventResponse } from './pocketbase-types';
    import { apiClient } from '$lib/gin';

    // AttendeeResponse but with .exapand
    export let invite;

    const dispatch = createEventDispatcher();

    function onClick() {
        // dispatch('message', {});
    }
    

    afterUpdate(() => {
        console.log(invite)
        apiClient.patch("/invites/"+invite.id,{paid:invite.paid})
    });
    
</script>

<div class="mb-2 collapse  border-base-300 bg-base-100 rounded-box">
    <div class="overflow-x-auto w-full">
        <table class="table table-compact w-full">
            <thead>
                <tr>
                    <th />
                    <th><p class="pl-2">Status</p></th>
                    <th>Paid</th>
                </tr>
            </thead>
            <tbody>
                <tr on:click={onClick} >
                    <td>
                        <div class="flex items-center space-x-3">
                            <div class="font-bold">
                                {invite.contact.first_name}
                                {invite.contact.last_name}
                            </div>
                        </div>
                    </td>
                    <td>
                        <span class="badge badge-lg">{invite.status}</span>
                    </td>
                    <td>
                        <label class="swap swap-flip text-3xl">
                            <!-- this hidden checkbox controls the state -->
                            <input bind:checked={invite.paid} on:click={onClick} type="checkbox" />

                            <div class="swap-on">💵</div>
                            <div class="swap-off">❌</div>
                        </label>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</div>
