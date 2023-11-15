
# Shell Bud

**Shell Bud** is a command-line utility tool designed to enhance productivity and streamline management of self-hosted services and machines. It provides an easy-to-use interface for managing remote machines, executing commands, and automating routine tasks.

## Features

- [x] **Machine Management**: Easily add, list, and remove machines, keeping track of their IP addresses and online status.
- [x] **Command Management**: Manage and organize frequently used commands with options to add, list, and remove them.
- [ ] **Macro System**: Pin commands for quick execution, including the option to set aliases for ease of use.
- [ ] **Remote Execution**: Execute commands on remote machines, either specified per command or set as a default for each macro.
- [ ] **Enhanced Visuals**: Color-coded output and emoji/Nerd Font icons for an improved user experience.
- [ ] **Background Command Execution**: Run commands in the background, allowing for multitasking and checking results later.
- [ ] **Template Commands**: Create complex commands with placeholders for arguments, enabling flexible and dynamic command execution.
- [ ] **Plugin**: Extend Shell Bud's functionality with plugins, allowing for custom commands and features.
 
## Getting Started
1. Clone the repo
2. `touch ~/.config/sb/config.yaml`
3. Install locally: `just install-local`
4. Test Shell Bud with:`sb`

## Usage

### Machine Management
Adding a Machine `sb machine add [name] [ip]`

Listing Machines `sb machine list`

Removing a Machine `sb machine remove [name]`

### Command Management
Adding a Command `sb commands add [name] [description] [command]`
List commands `sb commands list`
Remove a Command `sb commands remove [name]`

## Contributing

Contributions to Shell Bud are welcome! Please read our contributing guidelines for details on how to contribute to this project.

---

## Work in Progress (WIP)
Macro System (WIP)
Pinning a Command

`shell-bud commands pin [name] [alias] [defaultMachine]`

Remote Execution (WIP)

    Feature description and command usage instructions

Enhanced Visuals (WIP)

    Feature description and command usage instructions

Background Command Execution (WIP)

    Feature description and command usage instructions

Template Commands (WIP)

    Feature description and command usage instructions

