local function notify_error(message)
    vim.notify("ox: " .. message, vim.log.levels.ERROR)
end

local function preview()
    local lines = vim.api.nvim_buf_get_lines(0, 0, -1, false)
    local markdown = table.concat(lines, "\n") .. "\n"
    local binary = vim.g.ox_binary or "ox"

    local ok, err = pcall(vim.system, { binary }, {
        stdin = markdown,
        text = true,
    }, function(result)
        vim.schedule(function()
            if result.code ~= 0 then
                local message = vim.trim(result.stderr or "")
                if message == "" then
                    message = "renderer exited with status " .. result.code
                end
                notify_error(message)
                return
            end

            local path = vim.fn.tempname() .. ".html"
            local file, open_err = io.open(path, "w")
            if not file then
                notify_error("could not create preview: " .. open_err)
                return
            end

            local write_ok, write_err = file:write(result.stdout)
            local close_ok, close_err = file:close()
            if not write_ok then
                notify_error("could not write preview: " .. write_err)
                return
            end
            if not close_ok then
                notify_error("could not close preview: " .. close_err)
                return
            end

            local call_ok, command, browser_err = pcall(vim.ui.open, path)
            if not call_ok then
                notify_error("could not open browser: " .. tostring(command))
            elseif not command then
                notify_error("could not open browser: " .. (browser_err or "no opener available"))
            end
        end)
    end)

    if not ok then
        notify_error("could not start renderer: " .. tostring(err))
    end
end

vim.api.nvim_create_user_command("Ox", preview, {
    desc = "Preview the current buffer as Markdown in a browser",
})
