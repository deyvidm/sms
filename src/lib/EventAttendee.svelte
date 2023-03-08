<script lang="ts">
    import { createEventDispatcher, onMount } from 'svelte';
    import { attendees, get50Attendees } from './pocketbase';
    import type { AttendeeResponse, EventResponse } from './pocketbase-types';

    export let who;

    const dispatch = createEventDispatcher();

    function onClick() {
        // dispatch('message', {});
    }
</script>

<div class="left-5 mb-2 collapse  border-base-300 bg-base-100 rounded-box">
    <div class="overflow-x-auto w-full">
        <table class="table table-compact w-full">
            <thead>
                <tr>
                    <th></th>
                    <th><p class="pl-2">Status</p></th>
                    <th>Paid</th>
                </tr>
            </thead>
            <tbody>
                <tr>
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
                        <!-- <label class="label cursor-pointer">
                        <span class="label-text">paid</span> 
                        <input type="checkbox" bind:checked={who.paid} class="checkbox" />
                    </label> -->

                    <label class="swap swap-flip text-3xl">
  
                        <!-- this hidden checkbox controls the state -->
                        <input bind:value={who.paid} type="checkbox" />
                        
                        <div class="swap-on">💵</div>
                        <div class="swap-off">🫰</div>
                      </label>
                    </td>
                    <td></td>
                </tr>
            </tbody>
        </table>
    </div>
</div>
