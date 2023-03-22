package command

import (
	"fmt"
	"github.com/no-mole/venus/proto/pbuser"
	"github.com/spf13/cobra"
)

func NewUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user <subcommand>",
		Short: "user commands",
	}
	cmd.AddCommand(NewUserRegisterCommand())
	cmd.AddCommand(NewUserUnregisterCommand())
	cmd.AddCommand(NewUserLoginCommand())
	cmd.AddCommand(NewUserChangeStatusCommand())
	cmd.AddCommand(NewUserListCommand())
	cmd.AddCommand(NewUserNamespaceListCommand())
	cmd.AddCommand(NewUserChangePasswordCommand())
	cmd.AddCommand(NewUserResetPasswordCommand())
	return cmd
}

var (
	uid         string
	name        string
	password    string
	userStatus  int32
	newPassword string
)

func NewUserRegisterCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register <uid> <password> [options]",
		Short: "register user",
		Run:   userRegisterCommandFunc,
	}
	cmd.Flags().StringVar(&uid, "uid", "", "uid for register")
	cmd.Flags().StringVar(&name, "name", "", "name for register")
	cmd.Flags().StringVar(&password, "password", "", "password for register")
	return cmd
}

func userRegisterCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	info, err := client.UserRegister(cmd.Context(), uid, name, password)
	if err != nil {
		panic(err)
	}
	println(info.String())
}

func NewUserUnregisterCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unregister <uid> [options]",
		Short: "unregister user",
		Run:   userUnRegisterCommandFunc,
	}
	cmd.Flags().StringVar(&uid, "uid", "", "uid for unregister")
	return cmd
}

func userUnRegisterCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	info, err := client.UserUnregister(cmd.Context(), uid)
	if err != nil {
		panic(err)
	}
	println(info.String())
}

func NewUserLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login <uid> <password> [options]",
		Short: "user login",
		Run:   userLoginCommandFunc,
	}
	cmd.Flags().StringVar(&uid, "uid", "", "uid for login")
	cmd.Flags().StringVar(&password, "password", "", "password for login")
	return cmd
}

func userLoginCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.UserLogin(cmd.Context(), uid, password)
	if err != nil {
		panic(err)
	}
	println(resp.String())
}

func NewUserChangeStatusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "changeStatus <uid> <status> [options]",
		Short: "change user status",
		Run:   userChangeStatusCommandFunc,
	}
	cmd.Flags().StringVar(&uid, "uid", "", "uid for change status")
	cmd.Flags().Int32Var(&userStatus, "status", 0, "status for change status")
	return cmd
}

func userChangeStatusCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	err = client.UserChangeStatus(cmd.Context(), uid, pbuser.UserStatus(userStatus))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s status has changed", uid)
}

func NewUserListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [options]",
		Short: "user list",
		Run:   userListCommandFunc,
	}
	return cmd
}

func userListCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.UserList(cmd.Context())
	if err != nil {
		panic(err)
	}
	println(resp.String())
}

func NewUserNamespaceListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "namespace list <uid> [options]",
		Short: "user namespace list",
		Run:   userNamespaceListCommandFunc,
	}
	cmd.Flags().StringVar(&uid, "uid", "", "uid for user namespace list")
	return cmd
}

func userNamespaceListCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.UserNamespaceList(cmd.Context(), uid)
	if err != nil {
		panic(err)
	}
	println(resp.String())
}

func NewUserChangePasswordCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "changePassword <uid> <old-password> <new-password> [options]",
		Short: "user change password",
		Run:   userChangePasswordCommandFunc,
	}
	cmd.Flags().StringVar(&uid, "uid", "", "uid for user change password")
	cmd.Flags().StringVar(&password, "old-password", "", "password for user change password")
	cmd.Flags().StringVar(&newPassword, "new-password", "", "new-password for user change password")
	return cmd
}

func userChangePasswordCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.UserChangePassword(cmd.Context(), uid, password, newPassword)
	if err != nil {
		panic(err)
	}
	println(resp.String())
}

func NewUserResetPasswordCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reset password <uid> [options]",
		Short: "user reset password",
		Run:   userResetPasswordCommandFunc,
	}
	cmd.Flags().StringVar(&uid, "uid", "", "uid for user reset password")
	return cmd
}

func userResetPasswordCommandFunc(cmd *cobra.Command, _ []string) {
	var err error
	client, err := getClientFromFlags()
	if err != nil {
		panic(err)
	}
	resp, err := client.UserResetPassword(cmd.Context(), uid)
	if err != nil {
		panic(err)
	}
	println(resp.String())
}
