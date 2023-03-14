<script lang="ts">
    import { afterUpdate, createEventDispatcher, onMount } from 'svelte';
    import { attendees, get50Attendees, pb, updateAttendee } from './pocketbase';
    import type { AttendeeResponse, EventResponse } from './pocketbase-types';

    // AttendeeResponse but with .exapand
    export let who;

    const dispatch = createEventDispatcher();

    function onClick() {
        // dispatch('message', {});
    }
    

    afterUpdate(() => {
        updateAttendee(who)
            // .then((result) => console.log(result))
            .catch((err) => console.log(err));
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
                                {who.expand.contact.first_name}
                                {who.expand.contact.last_name}
                            </div>
                        </div>
                    </td>
                    <td>
                        <span class="badge badge-lg">{who.status}</span>
                    </td>
                    <td>
                        <label class="swap swap-flip text-3xl">
                            <!-- this hidden checkbox controls the state -->
                            <input bind:checked={who.paid} on:click={onClick} type="checkbox" />

                            <div class="swap-on">💵</div>
                            <div class="swap-off">🫰</div>
                        </label>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</div>
