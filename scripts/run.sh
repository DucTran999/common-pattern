#!/bin/bash

# path Variable
CONCURRENCY_PKG=concurrency

show_menu() {
    echo "========== MENU =========="
    echo "====== Concurrency ======="
    echo "1) worker pool"
    echo "2) generator"
    echo "3) multiplexing"
    echo "0) Exit"
    echo "=========================="
}

show_menu
read -p "Choose an option: " choice

case "$choice" in
1)
    clear
    echo "======== APP: Concurrency ========="
    echo "=========== Worker Pool ==========="
    go run $CONCURRENCY_PKG/main.go --app-name=worker-pool
    ;;
2)
    clear
    echo "======== APP: Concurrency ========="
    echo "============ Generator ============"
    go run $CONCURRENCY_PKG/main.go --app-name=generator
    ;;
3)
    clear
    echo "======== APP: Concurrency ========="
    echo "======== Multiplexing ========="
    go run $CONCURRENCY_PKG/main.go --app-name=multiplexing
    ;;
0)
    echo "Goodbye!"
    exit 0
    ;;
*)
    echo "Invalid option. Please try again."
    ;;
esac
