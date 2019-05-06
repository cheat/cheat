# Submit a new job:
sbatch job.sh

# List all jobs for a user:
squeue -u user_name

# Cancel a job by id or name:
scancel job_id
scancel --name job_name

# List all information for a job:
scontrol show jobid -dd job_id

# Status info for currently running job:
sstat --format=AveCPU,AvePages,AveRSS,AveVMSize,JobID -j job_id --allsteps
