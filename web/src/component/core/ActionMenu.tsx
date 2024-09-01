"use client";
import {Menu} from "antd";
import {useRouter} from "next/navigation";

export default function ActionMenu () {
    const handleLogout = () => {
        localStorage.removeItem("token");
        //reload page
        window.location.reload();
    }
    return (
        <Menu>
            <Menu.Item key="logout" onClick={handleLogout}>
                Đăng xuất
            </Menu.Item>
        </Menu>
    )
}